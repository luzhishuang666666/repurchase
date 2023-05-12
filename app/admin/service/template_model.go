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

type TemplateModel struct {
	service.Service
}

// GetPage 获取TemplateModel列表
func (e *TemplateModel) GetPage(c *dto.TemplateModelGetPageReq, p *actions.DataPermission, list *[]models.TemplateModel, count *int64) error {
	var err error
	var data models.TemplateModel

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("TemplateModelService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取TemplateModel对象
func (e *TemplateModel) Get(d *dto.TemplateModelGetReq, p *actions.DataPermission, model *models.TemplateModel) error {
	var data models.TemplateModel

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetTemplateModel error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建TemplateModel对象
func (e *TemplateModel) Insert(c *dto.TemplateModelInsertReq) error {
	var err error
	var data models.TemplateModel
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("TemplateModelService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改TemplateModel对象
func (e *TemplateModel) Update(c *dto.TemplateModelUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.TemplateModel{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("TemplateModelService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除TemplateModel
func (e *TemplateModel) Remove(d *dto.TemplateModelDeleteReq, p *actions.DataPermission) error {
	var data models.TemplateModel

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveTemplateModel error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
