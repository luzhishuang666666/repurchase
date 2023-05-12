package models

import (
	_ "time"

	"go-admin/common/models"
)

type RepurchaseInfo struct {
	models.Model

	ModelId  string `json:"modelId" gorm:"type:int unsigned;comment:模型编码"`
	RecordId string `json:"recordId" gorm:"type:int unsigned;comment:用户记录编码"`
	Status   string `json:"status" gorm:"type:int;comment:预测记录状态"`
	Result   string `json:"result" gorm:"type:int;comment:预测结果"`
	models.ModelTime
	models.ControlBy
}

func (RepurchaseInfo) TableName() string {
	return "repurchase_info"
}

func (e *RepurchaseInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RepurchaseInfo) GetId() interface{} {
	return e.Id
}
