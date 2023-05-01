package models

import (
	_ "time"

	"go-admin/common/models"
)

type Commodity struct {
	models.Model

	CommodityName       string `json:"commodityName" gorm:"type:varchar(128);comment:商品名称"`
	CommodityBrandId    string `json:"commodityBrandId" gorm:"type:int unsigned;comment:商品品牌id"`
	CommodityCategoryId string `json:"commodityCategoryId" gorm:"type:int unsigned;comment:商品品类id"`
	Avatar              string `json:"avatar" gorm:"type:varchar(256);comment:商品展示"`
	CommodityRemark     string `json:"commodityRemark" gorm:"type:varchar(256);comment:商品备注"`
	Status              string `json:"status" gorm:"type:int;comment:状态"`
	models.ModelTime
	models.ControlBy
}

func (Commodity) TableName() string {
	return "commodity"
}

func (e *Commodity) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Commodity) GetId() interface{} {
	return e.Id
}
