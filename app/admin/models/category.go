package models

import (
	_ "time"

	"go-admin/common/models"
)

type Category struct {
	models.Model

	CategoryName   string `json:"categoryName" gorm:"type:varchar(256);comment:品类名称"`
	CategoryRemark string `json:"categoryRemark" gorm:"type:varchar(1024);comment:品类备注"`
	Status         string `json:"status" gorm:"type:int;comment:状态"`
	models.ModelTime
	models.ControlBy
}

func (Category) TableName() string {
	return "category"
}

func (e *Category) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Category) GetId() interface{} {
	return e.Id
}
