package dto

import (
	_ "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type TemplateModelGetPageReq struct {
	dto.Pagination `search:"-"`
	TemplateName   string `form:"templateName"  search:"type:contains;column:template_name;table:template_model" comment:"模板名称"`
	TemplateType   string `form:"templateType"  search:"type:exact;column:template_type;table:template_model" comment:"模板类型"`
	TemplateModelOrder
}

type TemplateModelOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:template_model"`
	TemplateName  string `form:"templateNameOrder"  search:"type:order;column:template_name;table:template_model"`
	TemplateType  string `form:"templateTypeOrder"  search:"type:order;column:template_type;table:template_model"`
	TemplateParam string `form:"templateParamOrder"  search:"type:order;column:template_param;table:template_model"`
	CreatedAt     string `form:"createdAtOrder"  search:"type:order;column:created_at;table:template_model"`
	UpdatedAt     string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:template_model"`
	DeletedAt     string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:template_model"`
	CreateBy      string `form:"createByOrder"  search:"type:order;column:create_by;table:template_model"`
	UpdateBy      string `form:"updateByOrder"  search:"type:order;column:update_by;table:template_model"`
}

func (m *TemplateModelGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TemplateModelInsertReq struct {
	Id            int    `json:"-" comment:"编码"` // 编码
	TemplateName  string `json:"templateName" comment:"模板名称"`
	TemplateType  string `json:"templateType" comment:"模板类型"`
	TemplateParam string `json:"templateParam" comment:"模板参数"`
	common.ControlBy
}

func (s *TemplateModelInsertReq) Generate(model *models.TemplateModel) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.TemplateName = s.TemplateName
	model.TemplateType = s.TemplateType
	model.TemplateParam = s.TemplateParam
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *TemplateModelInsertReq) GetId() interface{} {
	return s.Id
}

type TemplateModelUpdateReq struct {
	Id            int    `uri:"id" comment:"编码"` // 编码
	TemplateName  string `json:"templateName" comment:"模板名称"`
	TemplateType  string `json:"templateType" comment:"模板类型"`
	TemplateParam string `json:"templateParam" comment:"模板参数"`
	common.ControlBy
}

func (s *TemplateModelUpdateReq) Generate(model *models.TemplateModel) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.TemplateName = s.TemplateName
	model.TemplateType = s.TemplateType
	model.TemplateParam = s.TemplateParam
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *TemplateModelUpdateReq) GetId() interface{} {
	return s.Id
}

// TemplateModelGetReq 功能获取请求参数
type TemplateModelGetReq struct {
	Id int `uri:"id"`
}

func (s *TemplateModelGetReq) GetId() interface{} {
	return s.Id
}

// TemplateModelDeleteReq 功能删除请求参数
type TemplateModelDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *TemplateModelDeleteReq) GetId() interface{} {
	return s.Ids
}
