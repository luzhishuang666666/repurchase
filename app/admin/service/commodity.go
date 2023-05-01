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

type Commodity struct {
	service.Service
}

// GetPage 获取Commodity列表
func (e *Commodity) GetPage(c *dto.CommodityGetPageReq, p *actions.DataPermission, list *[]models.Commodity, count *int64) error {
	var err error
	var data models.Commodity

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("CommodityService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Commodity对象
func (e *Commodity) Get(d *dto.CommodityGetReq, p *actions.DataPermission, model *models.Commodity) error {
	var data models.Commodity

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetCommodity error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Commodity对象
func (e *Commodity) Insert(c *dto.CommodityInsertReq) error {
	var err error
	var data models.Commodity
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("CommodityService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Commodity对象
func (e *Commodity) Update(c *dto.CommodityUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Commodity{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("CommodityService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Commodity
func (e *Commodity) Remove(d *dto.CommodityDeleteReq, p *actions.DataPermission) error {
	var data models.Commodity

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveCommodity error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
