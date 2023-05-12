package models

import (
	_ "time"

	"go-admin/common/models"
)

type TemplateModel struct {
	models.Model

	TemplateName  string `json:"templateName" gorm:"type:varchar(64);comment:模板名称"`
	TemplateType  string `json:"templateType" gorm:"type:varchar(64);comment:模板类型"`
	TemplateParam string `json:"templateParam" gorm:"type:longblob;comment:模板参数"`
	models.ModelTime
	models.ControlBy
}

func (TemplateModel) TableName() string {
	return "template_model"
}

func (e *TemplateModel) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TemplateModel) GetId() interface{} {
	return e.Id
}
