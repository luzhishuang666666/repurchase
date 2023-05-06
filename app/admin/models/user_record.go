package models

import (
	_ "time"

	"go-admin/common/models"
)

type UserRecord struct {
	models.Model

	ItemId      string `json:"itemId" gorm:"type:int unsigned;comment:购买者编号"`
	ShopId      string `json:"shopId" gorm:"type:int unsigned;comment:商店编号"`
	AgeRange    string `json:"ageRange" gorm:"type:int;comment:购买者年龄范围"`
	Gender      string `json:"gender" gorm:"type:int;comment:购买者性别"`
	CommodityId string `json:"commodityId" gorm:"type:int unsigned;comment:商品编号"`
	ActionType  string `json:"actionType" gorm:"type:int;comment:操作行为类别"`
	models.ModelTime
	models.ControlBy
}

func (UserRecord) TableName() string {
	return "user_record"
}

func (e *UserRecord) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *UserRecord) GetId() interface{} {
	return e.Id
}
