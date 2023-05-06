package models

import (
	_ "time"

	"go-admin/common/models"
)

type Shop struct {
	models.Model

	ShopNo     string `json:"shopNo" gorm:"type:varchar(32);comment:商店编号"`
	ShopName   string `json:"shopName" gorm:"type:varchar(64);comment:商店名称"`
	ShopRemark string `json:"shopRemark" gorm:"type:varchar(4096);comment:商店备注"`
	Status     string `json:"status" gorm:"type:int;comment:状态"`
	models.ModelTime
	models.ControlBy
}

func (Shop) TableName() string {
	return "shop"
}

func (e *Shop) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Shop) GetId() interface{} {
	return e.Id
}

func (e *Shop) GetStatus() interface{} {
	return e.Status
}

func (e *Shop) SetStatus(status string) {
	e.Status = status
}
