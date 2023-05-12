package dto

import (
	_ "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RepurchaseInfoGetPageReq struct {
	dto.Pagination `search:"-"`
	ModelId        string `form:"modelId"  search:"type:exact;column:model_id;table:repurchase_info" comment:"模型编码"`
	RecordId       string `form:"recordId"  search:"type:exact;column:record_id;table:repurchase_info" comment:"用户记录编码"`
	Status         string `form:"status"  search:"type:exact;column:status;table:repurchase_info" comment:"预测记录状态"`
	Result         string `form:"result"  search:"type:exact;column:result;table:repurchase_info" comment:"预测结果"`
	RepurchaseInfoOrder
}

type RepurchaseInfoOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:repurchase_info"`
	ModelId   string `form:"modelIdOrder"  search:"type:order;column:model_id;table:repurchase_info"`
	RecordId  string `form:"recordIdOrder"  search:"type:order;column:record_id;table:repurchase_info"`
	Status    string `form:"statusOrder"  search:"type:order;column:status;table:repurchase_info"`
	Result    string `form:"resultOrder"  search:"type:order;column:result;table:repurchase_info"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:repurchase_info"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:repurchase_info"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:repurchase_info"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:repurchase_info"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:repurchase_info"`
}

func (m *RepurchaseInfoGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RepurchaseInfoInsertReq struct {
	Id       int    `json:"-" comment:"编码"` // 编码
	ModelId  string `json:"modelId" comment:"模型编码"`
	RecordId string `json:"recordId" comment:"用户记录编码"`
	Status   string `json:"status" comment:"预测记录状态"`
	Result   string `json:"result" comment:"预测结果"`
	common.ControlBy
}

func (s *RepurchaseInfoInsertReq) Generate(model *models.RepurchaseInfo) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ModelId = s.ModelId
	model.RecordId = s.RecordId
	model.Status = s.Status
	model.Result = s.Result
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *RepurchaseInfoInsertReq) GetId() interface{} {
	return s.Id
}

type RepurchaseInfoUpdateReq struct {
	Id       int    `uri:"id" comment:"编码"` // 编码
	ModelId  string `json:"modelId" comment:"模型编码"`
	RecordId string `json:"recordId" comment:"用户记录编码"`
	Status   string `json:"status" comment:"预测记录状态"`
	Result   string `json:"result" comment:"预测结果"`
	common.ControlBy
}

func (s *RepurchaseInfoUpdateReq) Generate(model *models.RepurchaseInfo) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ModelId = s.ModelId
	model.RecordId = s.RecordId
	model.Status = s.Status
	model.Result = s.Result
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *RepurchaseInfoUpdateReq) GetId() interface{} {
	return s.Id
}

// RepurchaseInfoGetReq 功能获取请求参数
type RepurchaseInfoGetReq struct {
	Id int `uri:"id"`
}

func (s *RepurchaseInfoGetReq) GetId() interface{} {
	return s.Id
}

// RepurchaseInfoDeleteReq 功能删除请求参数
type RepurchaseInfoDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RepurchaseInfoDeleteReq) GetId() interface{} {
	return s.Ids
}
