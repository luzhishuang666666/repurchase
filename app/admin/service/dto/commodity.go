package dto

import (
	_ "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type CommodityGetPageReq struct {
	dto.Pagination      `search:"-"`
	CommodityName       string `form:"commodityName"  search:"type:exact;column:commodity_name;table:commodity" comment:"商品名称"`
	CommodityBrandId    string `form:"commodityBrandId"  search:"type:exact;column:commodity_brand_id;table:commodity" comment:"商品品牌id"`
	CommodityCategoryId string `form:"commodityCategoryId"  search:"type:exact;column:commodity_category_id;table:commodity" comment:"商品品类id"`
	CommodityOrder
}

type CommodityOrder struct {
	Id                  string `form:"idOrder"  search:"type:order;column:id;table:commodity"`
	CommodityName       string `form:"commodityNameOrder"  search:"type:order;column:commodity_name;table:commodity"`
	CommodityBrandId    string `form:"commodityBrandIdOrder"  search:"type:order;column:commodity_brand_id;table:commodity"`
	CommodityCategoryId string `form:"commodityCategoryIdOrder"  search:"type:order;column:commodity_category_id;table:commodity"`
	Avatar              string `form:"avatarOrder"  search:"type:order;column:avatar;table:commodity"`
	CommodityRemark     string `form:"commodityRemarkOrder"  search:"type:order;column:commodity_remark;table:commodity"`
	Status              string `form:"statusOrder"  search:"type:order;column:status;table:commodity"`
	CreatedAt           string `form:"createdAtOrder"  search:"type:order;column:created_at;table:commodity"`
	UpdatedAt           string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:commodity"`
	DeletedAt           string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:commodity"`
	CreateBy            string `form:"createByOrder"  search:"type:order;column:create_by;table:commodity"`
	UpdateBy            string `form:"updateByOrder"  search:"type:order;column:update_by;table:commodity"`
}

func (m *CommodityGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type CommodityInsertReq struct {
	Id                  int    `json:"-" comment:"编码"` // 编码
	CommodityName       string `json:"commodityName" comment:"商品名称"`
	CommodityBrandId    string `json:"commodityBrandId" comment:"商品品牌id"`
	CommodityCategoryId string `json:"commodityCategoryId" comment:"商品品类id"`
	Avatar              string `json:"avatar" comment:"商品展示"`
	CommodityRemark     string `json:"commodityRemark" comment:"商品备注"`
	Status              string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *CommodityInsertReq) Generate(model *models.Commodity) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CommodityName = s.CommodityName
	model.CommodityBrandId = s.CommodityBrandId
	model.CommodityCategoryId = s.CommodityCategoryId
	model.Avatar = s.Avatar
	model.CommodityRemark = s.CommodityRemark
	model.Status = s.Status
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *CommodityInsertReq) GetId() interface{} {
	return s.Id
}

type CommodityUpdateReq struct {
	Id                  int    `uri:"id" comment:"编码"` // 编码
	CommodityName       string `json:"commodityName" comment:"商品名称"`
	CommodityBrandId    string `json:"commodityBrandId" comment:"商品品牌id"`
	CommodityCategoryId string `json:"commodityCategoryId" comment:"商品品类id"`
	Avatar              string `json:"avatar" comment:"商品展示"`
	CommodityRemark     string `json:"commodityRemark" comment:"商品备注"`
	Status              string `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *CommodityUpdateReq) Generate(model *models.Commodity) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CommodityName = s.CommodityName
	model.CommodityBrandId = s.CommodityBrandId
	model.CommodityCategoryId = s.CommodityCategoryId
	model.Avatar = s.Avatar
	model.CommodityRemark = s.CommodityRemark
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *CommodityUpdateReq) GetId() interface{} {
	return s.Id
}

// CommodityGetReq 功能获取请求参数
type CommodityGetReq struct {
	Id int `uri:"id"`
}

func (s *CommodityGetReq) GetId() interface{} {
	return s.Id
}

// CommodityDeleteReq 功能删除请求参数
type CommodityDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *CommodityDeleteReq) GetId() interface{} {
	return s.Ids
}
