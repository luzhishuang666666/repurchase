package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-admin/common/storage"
	"sort"
	"strconv"
	_ "strconv"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type Shop struct {
	service.Service
}

// GetPage 获取Shop列表
func (e *Shop) GetPage(c *dto.ShopGetPageReq, p *actions.DataPermission, list *[]models.Shop, count *int64) error {
	var err error
	var data models.Shop

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Where("create_by = ?", c.CreateBy).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ShopService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Shop对象
func (e *Shop) Get(d *dto.ShopGetReq, p *actions.DataPermission, model *models.Shop) error {
	var data models.Shop

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetShop error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Shop对象
func (e *Shop) Insert(c *dto.ShopInsertReq) error {
	var err error
	var data models.Shop
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ShopService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Shop对象
func (e *Shop) Update(c *dto.ShopUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Shop{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("ShopService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Shop
func (e *Shop) Remove(d *dto.ShopDeleteReq, p *actions.DataPermission) error {
	var data models.Shop

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveShop error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

func (e *Shop) ChangeStatus(s *dto.ShopChangeStatusReq) error {

	var data models.Shop

	err := e.Orm.Model(&data).
		Where("id = ?", s.GetId()).First(&data).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	if data.GetStatus() == "1" {
		data.SetStatus("2")
	} else {
		data.SetStatus("1")
	}

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("ShopService Save error:%s \r\n", err)
		return err
	}

	return nil
}

type CommodityResult []struct {
	CommodityID int `json:"commodity_id"`
	Count       int `json:"count"`
}

func (e *Shop) ShopRank(r *dto.ShopRankReq, rank *dto.ShopRankResp) error {
	var UserRecordDatas []models.UserRecord
	var CommodityRes CommodityResult
	var Commodities []models.Commodity
	var CommoditiyIds []int
	var DayRank []dto.RankItem
	var MonthRank []dto.RankItem
	var YearRank []dto.RankItem
	id := r.GetId()

	key := "Rank_" + strconv.Itoa(id)
	// Redis 缓存是否存在

	RedisRank, RedisErr := storage.CacheAdapter.Get(key)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
			// Day
			fromDay := time.Now().Format("2006-01-02") + " 00:00:00"
			toDay := time.Now().Format("2006-01-02") + " 23:59:59"
			err := e.Orm.Model(&UserRecordDatas).
				Where("shop_id = ? AND action_type = ? AND created_at > ? AND created_at < ?",
					id, 2, fromDay, toDay).
				Order("COUNT(commodity_id) DESC").
				Select("commodity_id, count(*) as count").
				Group("commodity_id").
				Limit(5).
				Find(&CommodityRes).Error

			if err != nil {
				e.Log.Errorf("ShopService Save error:%s \r\n", err)
			}
			CommodityMap := make(map[int]models.Commodity)
			for _, item := range CommodityRes {
				CommoditiyIds = append(CommoditiyIds, item.CommodityID)
			}
			err = e.Orm.Model(&Commodities).
				Where("id IN ?", CommoditiyIds).
				Find(&Commodities).
				Error

			for _, c := range Commodities {
				if _, ok := CommodityMap[c.Id]; !ok {
					CommodityMap[c.Id] = c
				}
			}
			count := 0
			for _, s := range CommodityRes {
				var item dto.RankItem
				if count < 3 {
					item.CommodityTrend = 1
				}
				item.CommodityName = CommodityMap[s.CommodityID].CommodityName
				item.CommodityScore = s.Count
				DayRank = append(DayRank, item)
				count++
			}

			rank.Day = DayRank

			// Month
			fromMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location()).Format("2006-01-02 15:04:05")
			nextMonth := time.Date(time.Now().Year(), time.Now().Month()+1, 1, 0, 0, 0, 0, time.Now().Location())
			toMonth := nextMonth.Add(-time.Second).Format("2006-01-02 15:04:05")

			err = e.Orm.Model(&UserRecordDatas).
				Where("shop_id = ? AND action_type = ? AND created_at > ? AND created_at < ?",
					id, 2, fromMonth, toMonth).
				Order("COUNT(commodity_id) DESC").
				Select("commodity_id, count(*) as count").
				Group("commodity_id").
				Limit(5).
				Find(&CommodityRes).Error

			if err != nil {
				e.Log.Errorf("ShopService Save error:%s \r\n", err)
			}
			CommodityMap = make(map[int]models.Commodity)
			copy(CommoditiyIds[:], make([]int, len(CommoditiyIds)))
			for _, item := range CommodityRes {
				CommoditiyIds = append(CommoditiyIds, item.CommodityID)
			}
			err = e.Orm.Model(&Commodities).
				Where("id IN ?", CommoditiyIds).
				Find(&Commodities).
				Error

			for _, c := range Commodities {
				if _, ok := CommodityMap[c.Id]; !ok {
					CommodityMap[c.Id] = c
				}
			}

			count = 0
			for _, s := range CommodityRes {
				var item dto.RankItem
				if count < 3 {
					item.CommodityTrend = 1
				}
				item.CommodityName = CommodityMap[s.CommodityID].CommodityName
				item.CommodityScore = s.Count
				MonthRank = append(MonthRank, item)
				count++
			}

			rank.Month = MonthRank

			// year
			fromYear := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location()).Format("2006-01-02 15:04:05")
			nextYear := time.Date(time.Now().Year()+1, time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location())
			toYear := nextYear.Add(-time.Second).Format("2006-01-02 15:04:05")

			err = e.Orm.Model(&UserRecordDatas).
				Where("shop_id = ? AND action_type = ? AND created_at > ? AND created_at < ?",
					id, 2, fromYear, toYear).
				Order("COUNT(commodity_id) DESC").
				Select("commodity_id, count(*) as count").
				Group("commodity_id").
				Limit(5).
				Find(&CommodityRes).Error

			if err != nil {
				e.Log.Errorf("ShopService Save error:%s \r\n", err)
			}
			CommodityMap = make(map[int]models.Commodity)
			copy(CommoditiyIds[:], make([]int, len(CommoditiyIds)))
			for _, item := range CommodityRes {
				CommoditiyIds = append(CommoditiyIds, item.CommodityID)
			}
			err = e.Orm.Model(&Commodities).
				Where("id IN ?", CommoditiyIds).
				Find(&Commodities).
				Error

			for _, c := range Commodities {
				if _, ok := CommodityMap[c.Id]; !ok {
					CommodityMap[c.Id] = c
				}
			}

			count = 0
			for _, s := range CommodityRes {
				var item dto.RankItem
				if count < 3 {
					item.CommodityTrend = 1
				}
				item.CommodityName = CommodityMap[s.CommodityID].CommodityName
				item.CommodityScore = s.Count
				YearRank = append(YearRank, item)
				count++
			}

			rank.Year = YearRank

			rankString, err := json.Marshal(rank)
			if err != nil {
				e.Log.Errorf("json error", err)
			}
			now := time.Now()
			todayEnd := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
			seconds := int(todayEnd.Sub(now).Seconds())
			RedisErr = storage.CacheAdapter.Set(key, rankString, seconds)
			if RedisErr != nil {
				e.Log.Fatalf("cache error, %s\n", storage.Err.Error())
			}
		}
	}()

	if RedisErr != nil {
		e.Log.Fatalf("cache error, %s\n", storage.Err.Error())
		return RedisErr
	}
	if RedisRank != "" {
		err := json.Unmarshal([]byte(RedisRank), &rank)
		if err != nil {
			e.Log.Errorf("json error", err)
		}
	}
	return nil
}

func (e *Shop) ShopAnalise(r *dto.ShopAnaliseReq, analise *dto.ShopAnaliseResp) error {
	id := r.GetId()
	var UserRecordDatas []models.UserRecord
	var totalSales int
	// TotalSales
	fromDay := time.Now().Format("2006-01-02") + " 00:00:00"
	toDay := time.Now().Format("2006-01-02") + " 23:59:59"

	err := e.Orm.Model(&UserRecordDatas).
		Where("shop_id = ? AND action_type = ? AND created_at > ? AND created_at < ?",
			id, 2, fromDay, toDay).
		Select("count(*) as count").
		Find(&totalSales).Error

	if err != nil {
		e.Log.Errorf("ShopService query error:%s \r\n", err)
		return err
	}
	analise.TotalSales = strconv.Itoa(totalSales)

	// dayOnDayRate weekOnWeekRate
	today := time.Now().Format("2006-01-02")
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	yesterdayLast := time.Now().AddDate(0, 0, -1).Add(24 * time.Hour).Format("2006-01-02")
	lastWeek := time.Now().AddDate(0, 0, -7).Format("2006-01-02")

	var todayRecords, lastWeekRecords, lastDayRecords []models.UserRecord

	e.Orm.Model(&UserRecordDatas).Where("created_at >= ? AND created_at < ? AND action_type = ?", yesterday, today, 2).Find(&todayRecords)
	e.Orm.Model(&UserRecordDatas).Where("created_at >= ? AND created_at < ? AND shop_id = ? AND action_type = ?", lastWeek, yesterday, id, 2).Find(&lastWeekRecords)
	e.Orm.Model(&UserRecordDatas).Where("created_at >= ? AND created_at < ? AND shop_id = ? AND action_type = ?", yesterday, yesterdayLast, id, 2).Find(&lastDayRecords)

	saleCount := dto.SaleCount{}

	// 计算当天销量
	for _, record := range todayRecords {
		saleCount.TodayCount++
		record.GetId()
	}

	// 计算上一天同一时间段的销量
	for _, record := range lastDayRecords {
		saleCount.YesterdayCount++
		record.GetId()
	}

	// 计算上一周同一天同一时间段的销量
	for _, record := range lastWeekRecords {
		saleCount.LastWeekCount++
		record.GetId()
	}
	e.Log.Info("saleCount:", saleCount)
	var dayOnDayRate, weekOnWeekRate float64

	if saleCount.YesterdayCount > 0 {
		dayOnDayRate = float64(saleCount.TodayCount-saleCount.YesterdayCount) / float64(saleCount.YesterdayCount)
	}

	if saleCount.LastWeekCount > 0 {
		weekOnWeekRate = float64(saleCount.TodayCount-saleCount.LastWeekCount) / float64(saleCount.LastWeekCount)
	}

	analise.DayOnDayRate = strconv.FormatFloat(dayOnDayRate*100, 'f', 2, 64)
	analise.WeekOnWeekRate = strconv.FormatFloat(weekOnWeekRate*100, 'f', 2, 64)

	if dayOnDayRate > 0 {
		analise.DayTrend = 1
	} else {
		analise.DayTrend = 0
	}

	if weekOnWeekRate > 0 {
		analise.WeekTrend = 1
	} else {
		analise.WeekTrend = 0
	}

	// dailySales
	var salesRecords []models.UserRecord
	e.Orm.Model(&UserRecordDatas).Where("shop_id = ? AND action_type = ?", id, 2).Order("created_at asc").Find(&salesRecords)

	totalSales = 0
	var totalDays int
	for i, record := range salesRecords {
		// 第一条记录的创建时间为最早时间
		if i == 0 {
			totalDays = int(time.Now().Sub(record.CreatedAt).Hours() / 24)
		} else {
			days := int(record.CreatedAt.Sub(salesRecords[i-1].CreatedAt).Hours() / 24)
			totalDays += days
		}
		totalSales += 1
	}
	analise.DailySales = strconv.FormatFloat(float64(totalSales)/float64(totalDays), 'f', 2, 64)

	// totalViews
	var totalViews int64
	e.Orm.Model(&UserRecordDatas).Where("shop_id = ? and action_type = ?", id, 0).Count(&totalViews)
	analise.TotalViews = int(totalViews)

	// dailyViews
	var dailyViews int64
	e.Orm.Model(&UserRecordDatas).
		Where("shop_id = ? AND action_type = ? AND created_at >= ? AND created_at < ?", id, 0, yesterday, today).
		Count(&dailyViews)
	analise.DailyViews = int(dailyViews)

	// totalFavorites
	var totalFavorites int64
	e.Orm.Model(&UserRecordDatas).
		Where("shop_id = ? AND (action_type = 1 OR action_type = 3)", id).
		Count(&totalFavorites)
	analise.TotalFavorites = int(totalFavorites)

	// dailyFavorites
	var dailyFavorites int64
	e.Orm.Model(&UserRecordDatas).
		Where("shop_id = ? AND (action_type = 1 OR action_type = 3) AND created_at >= ? AND created_at < ?", id, yesterday, today).
		Count(&dailyFavorites)
	analise.DailyFavorites = int(dailyFavorites)

	// repurchaseRate
	var multiplePurchaseCount int64
	e.Orm.Model(&UserRecordDatas).
		Where("shop_id = ? AND action_type = ?", id, 2).
		Select("COUNT(CONCAT(commodity_id, item_id))").
		Count(&multiplePurchaseCount)

	// 统计购买次数为1的商品数量
	var singlePurchaseCount int64
	e.Orm.Model(&UserRecordDatas).
		Where("shop_id = ? AND action_type = ?", id, 2).
		Select("COUNT(DISTINCT CONCAT(commodity_id, item_id))").
		Count(&singlePurchaseCount)

	// 计算复购率
	analise.RepurchaseRate = int(float64(multiplePurchaseCount-singlePurchaseCount) / float64(multiplePurchaseCount) * 100)
	analise.MultiplePurchaseCount = int(multiplePurchaseCount)
	analise.SinglePurchaseCount = int(singlePurchaseCount)

	// barData
	var barData []dto.PurchaseCount
	e.Orm.Model(&UserRecordDatas).
		Select("DATE(created_at) AS x, COUNT(DISTINCT CONCAT(commodity_id, item_id)) AS y").
		Where("shop_id = ? AND action_type = 2 AND created_at >= DATE_SUB(CURDATE(), INTERVAL 7 DAY) AND deleted_at IS NULL", id).
		Group("DATE(created_at)").
		Scan(&barData)

	for i := range barData {
		t, err := time.Parse(time.RFC3339, barData[i].X)
		if err != nil {
			fmt.Println("解析时间字符串失败:", err)
			return err
		}
		barData[i].X = t.Format("2006-01-02")
	}

	e.Log.Info("barData:", barData)

	// 将不存在的日期补全为0
	dateMap := make(map[string]int)
	for _, record := range barData {
		dateMap[record.X] = record.Y
	}

	now := time.Now()
	for i := 6; i >= 0; i-- {
		date := now.AddDate(0, 0, -i).Format("2006-01-02")
		if _, ok := dateMap[date]; !ok {
			barData = append(barData, dto.PurchaseCount{X: date, Y: 0})
		}
	}

	// 按日期升序排序
	sort.Slice(barData, func(i, j int) bool {
		return barData[i].X < barData[j].X
	})
	analise.BarData = barData
	return nil
}
