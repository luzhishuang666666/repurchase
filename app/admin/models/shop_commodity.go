package models

import (
	_ "time"

	"go-admin/common/models"
)

type ShopCommodity struct {
	models.Model

	CommodityId string `json:"commodityId" gorm:"type:int unsigned;comment:商品编号"`
	ShopId      string `json:"shopId" gorm:"type:int unsigned;comment:商店编号"`
	Status      string `json:"status" gorm:"type:int;comment:状态"`
	models.ModelTime
	models.ControlBy
}

func (ShopCommodity) TableName() string {
	return "shop_commodity"
}

func (e *ShopCommodity) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ShopCommodity) GetId() interface{} {
	return e.Id
}
