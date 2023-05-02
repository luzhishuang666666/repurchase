package models

import (
	_ "time"

	"go-admin/common/models"
)

type Application struct {
	models.Model

	Type               string `json:"type" gorm:"type:int unsigned;comment:申请类型"`
	ApplicationContext string `json:"applicationContext" gorm:"type:varchar(8192);comment:申请内容"`
	ApplicationJson    string `json:"applicationJson" gorm:"type:varchar(4096);comment:原始申请"`
	Status             string `json:"status" gorm:"type:int;comment:状态"`
	models.ModelTime
	models.ControlBy
}

func (Application) TableName() string {
	return "application"
}

func (e *Application) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Application) GetId() interface{} {
	return e.Id
}

func (e *Application) SetStatus(status string) interface{} {
	e.Status = status
	return e.Status
}
