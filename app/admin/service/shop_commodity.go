package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type ShopCommodity struct {
	service.Service
}

// GetPage 获取ShopCommodity列表
func (e *ShopCommodity) GetPage(c *dto.ShopCommodityGetPageReq, p *actions.DataPermission, list *[]models.ShopCommodity, count *int64) error {
	var err error
	var data models.ShopCommodity

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ShopCommodityService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取ShopCommodity对象
func (e *ShopCommodity) Get(d *dto.ShopCommodityGetReq, p *actions.DataPermission, model *models.ShopCommodity) error {
	var data models.ShopCommodity

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetShopCommodity error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建ShopCommodity对象
func (e *ShopCommodity) Insert(c *dto.ShopCommodityInsertReq) error {
	var err error
	var data models.ShopCommodity
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ShopCommodityService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改ShopCommodity对象
func (e *ShopCommodity) Update(c *dto.ShopCommodityUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.ShopCommodity{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("ShopCommodityService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除ShopCommodity
func (e *ShopCommodity) Remove(d *dto.ShopCommodityDeleteReq, p *actions.DataPermission) error {
	var data models.ShopCommodity

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveShopCommodity error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
