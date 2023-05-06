package dto

import (
	_ "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ShopGetPageReq struct {
	dto.Pagination `search:"-"`
	ShopNo         string `form:"shopNo"  search:"type:contains;column:shop_no;table:shop" comment:"商店编号"`
	ShopName       string `form:"shopName"  search:"type:contains;column:shop_name;table:shop" comment:"商店名称"`
	ShopRemark     string `form:"shopRemark"  search:"type:contains;column:shop_remark;table:shop" comment:"商店备注"`
	Status         string `form:"status"  search:"type:exact;column:status;table:shop" comment:"状态"`
	ShopOrder
}

type ShopOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:shop"`
	ShopNo     string `form:"shopNoOrder"  search:"type:order;column:shop_no;table:shop"`
	ShopName   string `form:"shopNameOrder"  search:"type:order;column:shop_name;table:shop"`
	ShopRemark string `form:"shopRemarkOrder"  search:"type:order;column:shop_remark;table:shop"`
	Status     string `form:"statusOrder"  search:"type:order;column:status;table:shop"`
	CreatedAt  string `form:"createdAtOrder"  search:"type:order;column:created_at;table:shop"`
	UpdatedAt  string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:shop"`
	DeletedAt  string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:shop"`
	CreateBy   string `form:"createByOrder"  search:"type:order;column:create_by;table:shop"`
	UpdateBy   string `form:"updateByOrder"  search:"type:order;column:update_by;table:shop"`
}

func (m *ShopGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ShopInsertReq struct {
	Id         int    `json:"-" comment:"编码"` // 编码
	ShopNo     string `json:"shopNo" comment:"商店编号"`
	ShopName   string `json:"shopName" comment:"商店名称"`
	ShopRemark string `json:"shopRemark" comment:"商店备注"`
	Status     string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *ShopInsertReq) Generate(model *models.Shop) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ShopNo = s.ShopNo
	model.ShopName = s.ShopName
	model.ShopRemark = s.ShopRemark
	model.Status = s.Status
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *ShopInsertReq) GetId() interface{} {
	return s.Id
}

type ShopUpdateReq struct {
	Id         int    `uri:"id" comment:"编码"` // 编码
	ShopNo     string `json:"shopNo" comment:"商店编号"`
	ShopName   string `json:"shopName" comment:"商店名称"`
	ShopRemark string `json:"shopRemark" comment:"商店备注"`
	Status     string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *ShopUpdateReq) Generate(model *models.Shop) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ShopNo = s.ShopNo
	model.ShopName = s.ShopName
	model.ShopRemark = s.ShopRemark
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *ShopUpdateReq) GetId() interface{} {
	return s.Id
}

// ShopGetReq 功能获取请求参数
type ShopGetReq struct {
	Id int `uri:"id"`
}

func (s *ShopGetReq) GetId() interface{} {
	return s.Id
}

// ShopDeleteReq 功能删除请求参数
type ShopDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ShopDeleteReq) GetId() interface{} {
	return s.Ids
}

func (s *ShopInsertReq) SetShopNo(no string) {
	s.ShopNo = no
}

func (s *ShopInsertReq) SetStatus(status string) {
	s.Status = status
}

// ShopChangeStatusReq 功能获取请求参数
type ShopChangeStatusReq struct {
	Id int `uri:"id"`
}

func (s *ShopChangeStatusReq) GetId() int {
	return s.Id
}
