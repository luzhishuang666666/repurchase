package dto

import (
	_ "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ApplicationGetPageReq struct {
	dto.Pagination     `search:"-"`
	Type               string `form:"type"  search:"type:exact;column:type;table:application" comment:"申请类型"`
	ApplicationContext string `form:"applicationContext"  search:"type:contains;column:application_context;table:application" comment:"申请内容"`
	Status             string `form:"status"  search:"type:exact;column:status;table:application" comment:"状态"`
	ApplicationOrder
}

type ApplicationOrder struct {
	Id                 string `form:"idOrder"  search:"type:order;column:id;table:application"`
	Type               string `form:"typeOrder"  search:"type:order;column:type;table:application"`
	ApplicationContext string `form:"applicationContextOrder"  search:"type:order;column:application_context;table:application"`
	ApplicationJson    string `form:"applicationJsonOrder"  search:"type:order;column:application_json;table:application"`
	Status             string `form:"statusOrder"  search:"type:order;column:status;table:application"`
	CreatedAt          string `form:"createdAtOrder"  search:"type:order;column:created_at;table:application"`
	UpdatedAt          string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:application"`
	DeletedAt          string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:application"`
	CreateBy           string `form:"createByOrder"  search:"type:order;column:create_by;table:application"`
	UpdateBy           string `form:"updateByOrder"  search:"type:order;column:update_by;table:application"`
}

func (m *ApplicationGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ApplicationInsertReq struct {
	Id                 int    `json:"-" comment:"编码"` // 编码
	Type               string `json:"type" comment:"申请类型"`
	ApplicationContext string `json:"applicationContext" comment:"申请内容"`
	ApplicationJson    string `json:"applicationJson" comment:"原始申请"`
	Status             string `json:"status" comment:"状态"`
	common.ControlBy
}

type ApplicationApprovalReq struct {
	Id      int    `json:"id" comment:"编码"`
	Type    string `json:"type" comment:"申请类型"`
	Opinion int    `json:"opinion" comment:"审批意见"`
}

func (e *ApplicationApprovalReq) GetId() interface{} {
	return e.Id
}

func (e *ApplicationApprovalReq) GetOpinion() interface{} {
	return e.Opinion
}

func (s *ApplicationInsertReq) Generate(model *models.Application) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Type = s.Type
	model.ApplicationContext = s.ApplicationContext
	model.ApplicationJson = s.ApplicationJson
	model.Status = s.Status
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *ApplicationInsertReq) GetId() interface{} {
	return s.Id
}

type ApplicationUpdateReq struct {
	Id                 int    `uri:"id" comment:"编码"` // 编码
	Type               string `json:"type" comment:"申请类型"`
	ApplicationContext string `json:"applicationContext" comment:"申请内容"`
	ApplicationJson    string `json:"applicationJson" comment:"原始申请"`
	Status             string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *ApplicationUpdateReq) Generate(model *models.Application) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Type = s.Type
	model.ApplicationContext = s.ApplicationContext
	model.ApplicationJson = s.ApplicationJson
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *ApplicationUpdateReq) GetId() interface{} {
	return s.Id
}

// ApplicationGetReq 功能获取请求参数
type ApplicationGetReq struct {
	Id int `uri:"id"`
}

func (s *ApplicationGetReq) GetId() interface{} {
	return s.Id
}

// ApplicationDeleteReq 功能删除请求参数
type ApplicationDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ApplicationDeleteReq) GetId() interface{} {
	return s.Ids
}
