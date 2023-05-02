package service

import (
	"errors"
	log "github.com/go-admin-team/go-admin-core/logger"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type Application struct {
	service.Service
}

// GetPage 获取Application列表
func (e *Application) GetPage(c *dto.ApplicationGetPageReq, p *actions.DataPermission, list *[]models.Application, count *int64) error {
	var err error
	var data models.Application

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ApplicationService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Application对象
func (e *Application) Get(d *dto.ApplicationGetReq, p *actions.DataPermission, model *models.Application) error {
	var data models.Application

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetApplication error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Application对象
func (e *Application) Insert(c *dto.ApplicationInsertReq) error {
	var err error
	var data models.Application
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ApplicationService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Application对象
func (e *Application) Update(c *dto.ApplicationUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Application{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("ApplicationService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Application
func (e *Application) Remove(d *dto.ApplicationDeleteReq, p *actions.DataPermission) error {
	var data models.Application

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveApplication error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

func (e *Application) Approval(d *dto.ApplicationApprovalReq) error {
	var data models.Application

	err := e.Orm.Model(&data).First(d.GetId()).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	if d.GetOpinion() == 0 {
		data.SetStatus("2")
	} else {
		data.SetStatus("1")
	}

	log.Info("data :", data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("ApplicationService Save error:%s \r\n", err)
		return err
	}

	return nil
}
