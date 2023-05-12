package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/common/myKafka"
	"go-admin/common/storage"
	"gorm.io/gorm"
	"strconv"
	"time"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type ShopModel struct {
	service.Service
}

// GetPage 获取ShopModel列表
func (e *ShopModel) GetPage(c *dto.ShopModelGetPageReq, p *actions.DataPermission, list *[]models.ShopModel, count *int64) error {
	var err error
	var data models.ShopModel

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
		e.Log.Errorf("ShopModelService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取ShopModel对象
func (e *ShopModel) Get(d *dto.ShopModelGetReq, p *actions.DataPermission, model *models.ShopModel) error {
	var data models.ShopModel

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetShopModel error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建ShopModel对象
func (e *ShopModel) Insert(c *dto.ShopModelInsertReq) error {
	var err error
	var data models.ShopModel
	c.Generate(&data)

	var templateModel models.TemplateModel
	e.Orm.Model(&templateModel).
		Where("id = ?", data.TemplateId).
		Find(&templateModel)

	data.ModelParam = templateModel.TemplateParam
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ShopModelService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改ShopModel对象
func (e *ShopModel) Update(c *dto.ShopModelUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.ShopModel{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("ShopModelService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除ShopModel
func (e *ShopModel) Remove(d *dto.ShopModelDeleteReq, p *actions.DataPermission) error {
	var data models.ShopModel

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveShopModel error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

func (e *ShopModel) Forecast(d *dto.ShopModelForecastReq, p *actions.DataPermission) error {
	var forecastMq dto.ShopForecastMq
	var data models.ShopModel

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(&data, d.Id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetShopModel error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	forecastMq.ModelId = data.Id

	var templateModel models.TemplateModel
	err = e.Orm.Model(&templateModel).
		Scopes(
			actions.Permission(templateModel.TableName(), p),
		).
		First(&templateModel, data.TemplateId).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetShopModel error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	forecastMq.ModelType = templateModel.TemplateType

	var userRecordData models.UserRecord
	err = e.Orm.Model(&userRecordData).
		Scopes(
			actions.Permission(userRecordData.TableName(), p),
		).
		First(&userRecordData, d.UserRecordId).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetShopModel error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	forecastMq.ModelData.UserId = userRecordData.ItemId
	forecastMq.ModelData.MerchantId = userRecordData.ShopId
	forecastMq.ModelData.AgeRange = userRecordData.AgeRange
	forecastMq.ModelData.Gender = userRecordData.Gender

	var userRecordDatas []models.UserRecord
	err = e.Orm.Model(&userRecordDatas).
		Scopes(
			actions.Permission(userRecordData.TableName(), p),
		).
		Where("create_by = ?", userRecordData.CreateBy).
		Find(&userRecordDatas).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetShopModel error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	for _, record := range userRecordDatas {
		createdAt := record.CreatedAt
		actionType := record.ActionType

		if createdAt.Month() == time.November {
			switch actionType {
			case "0":
				forecastMq.ModelData.ClickInEleven++
			case "2":
				forecastMq.ModelData.BuyInEleven++
			case "3":
				forecastMq.ModelData.CollectionInEleven++
			}
		} else {
			switch actionType {
			case "0":
				forecastMq.ModelData.ClickNotInEleven++
			case "2":
				forecastMq.ModelData.BuyNotInEleven++
			case "3":
				forecastMq.ModelData.CollectionNotInEleven++
			}
		}
	}

	//保存预测记录
	var repurchaseInfoData models.RepurchaseInfo
	repurchaseInfoData.ModelId = strconv.Itoa(data.Id)
	repurchaseInfoData.RecordId = strconv.Itoa(userRecordData.Id)
	repurchaseInfoData.Status = "1"
	repurchaseInfoData.Result = ""
	repurchaseInfoData.SetCreateBy(data.CreateBy)

	err = e.Orm.Model(&repurchaseInfoData).Create(&repurchaseInfoData).Error
	if err != nil {
		e.Log.Fatalf("repurchase save error, %s\n", err)
		return err
	}

	e.Log.Info("111111111111111111111111111111")
	forecastMq.RepurchaseInfoId = repurchaseInfoData.Id
	e.Log.Info("2222222222222222222222222222222")
	// 发送mq消息
	MqMessage, err := json.Marshal(forecastMq)
	if err != nil {
		fmt.Println("转换为 JSON 字符串时出错:", err)
		return err
	}
	e.Log.Info("33333333333333333333333333333333")
	err = myKafka.SendMessage(string(MqMessage), "repurchasePredict", storage.KafkaConn)
	e.Log.Info("MqMessage:", MqMessage)
	if err != nil {
		e.Log.Errorf("kafka setup error, %s\n", err)
		return err
	}

	return nil
}
