package dto

import (
	_ "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type BrandGetPageReq struct {
	dto.Pagination `search:"-"`
	BrandName      string `form:"brandName"  search:"type:contains;column:brand_name;table:brand" comment:"品牌名称"`
	BrandRemark    string `form:"brandRemark"  search:"type:contains;column:brand_remark;table:brand" comment:"品牌备注"`
	BrandOrder
}

type BrandOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:brand"`
	BrandName   string `form:"brandNameOrder"  search:"type:order;column:brand_name;table:brand"`
	BrandRemark string `form:"brandRemarkOrder"  search:"type:order;column:brand_remark;table:brand"`
	Status      string `form:"statusOrder"  search:"type:order;column:status;table:brand"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:brand"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:brand"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:brand"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:brand"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:brand"`
}

func (m *BrandGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type BrandInsertReq struct {
	Id          int    `json:"-" comment:"编码"` // 编码
	BrandName   string `json:"brandName" comment:"品牌名称"`
	BrandRemark string `json:"brandRemark" comment:"品牌备注"`
	Status      string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *BrandInsertReq) Generate(model *models.Brand) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.BrandName = s.BrandName
	model.BrandRemark = s.BrandRemark
	model.Status = s.Status
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *BrandInsertReq) GetId() interface{} {
	return s.Id
}

type BrandUpdateReq struct {
	Id          int    `uri:"id" comment:"编码"` // 编码
	BrandName   string `json:"brandName" comment:"品牌名称"`
	BrandRemark string `json:"brandRemark" comment:"品牌备注"`
	Status      string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *BrandUpdateReq) Generate(model *models.Brand) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.BrandName = s.BrandName
	model.BrandRemark = s.BrandRemark
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *BrandUpdateReq) GetId() interface{} {
	return s.Id
}

// BrandGetReq 功能获取请求参数
type BrandGetReq struct {
	Id int `uri:"id"`
}

func (s *BrandGetReq) GetId() interface{} {
	return s.Id
}

// BrandDeleteReq 功能删除请求参数
type BrandDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *BrandDeleteReq) GetId() interface{} {
	return s.Ids
}
