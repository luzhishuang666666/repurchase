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

type RepurchaseInfo struct {
	service.Service
}

// GetPage 获取RepurchaseInfo列表
func (e *RepurchaseInfo) GetPage(c *dto.RepurchaseInfoGetPageReq, p *actions.DataPermission, list *[]models.RepurchaseInfo, count *int64) error {
	var err error
	var data models.RepurchaseInfo

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
		e.Log.Errorf("RepurchaseInfoService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RepurchaseInfo对象
func (e *RepurchaseInfo) Get(d *dto.RepurchaseInfoGetReq, p *actions.DataPermission, model *models.RepurchaseInfo) error {
	var data models.RepurchaseInfo

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRepurchaseInfo error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RepurchaseInfo对象
func (e *RepurchaseInfo) Insert(c *dto.RepurchaseInfoInsertReq) error {
	var err error
	var data models.RepurchaseInfo
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RepurchaseInfoService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RepurchaseInfo对象
func (e *RepurchaseInfo) Update(c *dto.RepurchaseInfoUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RepurchaseInfo{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RepurchaseInfoService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RepurchaseInfo
func (e *RepurchaseInfo) Remove(d *dto.RepurchaseInfoDeleteReq, p *actions.DataPermission) error {
	var data models.RepurchaseInfo

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRepurchaseInfo error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
