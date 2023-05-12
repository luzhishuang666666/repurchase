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

type Category struct {
	service.Service
}

// GetPage 获取Category列表
func (e *Category) GetPage(c *dto.CategoryGetPageReq, p *actions.DataPermission, list *[]models.Category, count *int64) error {
	var err error
	var data models.Category

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
		e.Log.Errorf("CategoryService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Category对象
func (e *Category) Get(d *dto.CategoryGetReq, p *actions.DataPermission, model *models.Category) error {
	var data models.Category

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetCategory error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Category对象
func (e *Category) Insert(c *dto.CategoryInsertReq) error {
	var err error
	var data models.Category
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("CategoryService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Category对象
func (e *Category) Update(c *dto.CategoryUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Category{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("CategoryService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Category
func (e *Category) Remove(d *dto.CategoryDeleteReq, p *actions.DataPermission) error {
	var data models.Category

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveCategory error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
