package models

import (
	_ "time"

	"go-admin/common/models"
)

type Brand struct {
	models.Model

	BrandName   string `json:"brandName" gorm:"type:varchar(128);comment:品牌名称"`
	BrandRemark string `json:"brandRemark" gorm:"type:varchar(256);comment:品牌备注"`
	Status      string `json:"status" gorm:"type:int;comment:状态"`
	models.ModelTime
	models.ControlBy
}

func (Brand) TableName() string {
	return "brand"
}

func (e *Brand) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Brand) GetId() interface{} {
	return e.Id
}
