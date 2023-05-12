package models

import (
	_ "time"

	"go-admin/common/models"
)

type ShopModel struct {
	models.Model

	ShopId      string `json:"shopId" gorm:"type:int unsigned;comment:商店编码"`
	TemplateId  string `json:"templateId" gorm:"type:int unsigned;comment:模板编码"`
	ModelName   string `json:"modelName" gorm:"type:varchar(64);comment:模型名称"`
	ModelRemark string `json:"modelRemark" gorm:"type:varchar(64);comment:模型备注"`
	ModelParam  string `json:"modelParam" gorm:"type:longblob;comment:模型参数"`
	models.ModelTime
	models.ControlBy
}

func (ShopModel) TableName() string {
	return "shop_model"
}

func (e *ShopModel) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ShopModel) GetId() interface{} {
	return e.Id
}
