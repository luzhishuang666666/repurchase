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

type UserRecord struct {
	service.Service
}

// GetPage 获取UserRecord列表
func (e *UserRecord) GetPage(c *dto.UserRecordGetPageReq, p *actions.DataPermission, list *[]models.UserRecord, count *int64) error {
	var err error
	var data models.UserRecord

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("UserRecordService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取UserRecord对象
func (e *UserRecord) Get(d *dto.UserRecordGetReq, p *actions.DataPermission, model *models.UserRecord) error {
	var data models.UserRecord

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetUserRecord error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建UserRecord对象
func (e *UserRecord) Insert(c *dto.UserRecordInsertReq) error {
	var err error
	var data models.UserRecord
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("UserRecordService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改UserRecord对象
func (e *UserRecord) Update(c *dto.UserRecordUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.UserRecord{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("UserRecordService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除UserRecord
func (e *UserRecord) Remove(d *dto.UserRecordDeleteReq, p *actions.DataPermission) error {
	var data models.UserRecord

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveUserRecord error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
