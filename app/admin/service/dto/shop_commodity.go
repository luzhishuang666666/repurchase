package dto

import (
	_ "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ShopCommodityGetPageReq struct {
	dto.Pagination `search:"-"`
	CommodityId    string `form:"commodityId"  search:"type:contains;column:commodity_id;table:shop_commodity" comment:"商品编号"`
	ShopId         string `form:"shopId"  search:"type:contains;column:shop_id;table:shop_commodity" comment:"商店编号"`
	Status         string `form:"status"  search:"type:exact;column:status;table:shop_commodity" comment:"状态"`
	ShopCommodityOrder
}

type ShopCommodityOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:shop_commodity"`
	CommodityId string `form:"commodityIdOrder"  search:"type:order;column:commodity_id;table:shop_commodity"`
	ShopId      string `form:"shopIdOrder"  search:"type:order;column:shop_id;table:shop_commodity"`
	Status      string `form:"statusOrder"  search:"type:order;column:status;table:shop_commodity"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:shop_commodity"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:shop_commodity"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:shop_commodity"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:shop_commodity"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:shop_commodity"`
}

func (m *ShopCommodityGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ShopCommodityInsertReq struct {
	Id          int    `json:"-" comment:"编码"` // 编码
	CommodityId string `json:"commodityId" comment:"商品编号"`
	ShopId      string `json:"shopId" comment:"商店编号"`
	Status      string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *ShopCommodityInsertReq) Generate(model *models.ShopCommodity) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CommodityId = s.CommodityId
	model.ShopId = s.ShopId
	model.Status = s.Status
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *ShopCommodityInsertReq) GetId() interface{} {
	return s.Id
}

type ShopCommodityUpdateReq struct {
	Id          int    `uri:"id" comment:"编码"` // 编码
	CommodityId string `json:"commodityId" comment:"商品编号"`
	ShopId      string `json:"shopId" comment:"商店编号"`
	Status      string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *ShopCommodityUpdateReq) Generate(model *models.ShopCommodity) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CommodityId = s.CommodityId
	model.ShopId = s.ShopId
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *ShopCommodityUpdateReq) GetId() interface{} {
	return s.Id
}

// ShopCommodityGetReq 功能获取请求参数
type ShopCommodityGetReq struct {
	Id int `uri:"id"`
}

func (s *ShopCommodityGetReq) GetId() interface{} {
	return s.Id
}

// ShopCommodityDeleteReq 功能删除请求参数
type ShopCommodityDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ShopCommodityDeleteReq) GetId() interface{} {
	return s.Ids
}
