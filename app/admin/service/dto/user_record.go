package dto

import (
	_ "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type UserRecordGetPageReq struct {
	dto.Pagination `search:"-"`
	ItemId         string `form:"itemId"  search:"type:exact;column:item_id;table:user_record" comment:"购买者编号"`
	ShopId         string `form:"shopId"  search:"type:contains;column:shop_id;table:user_record" comment:"商店编号"`
	AgeRange       string `form:"ageRange"  search:"type:exact;column:age_range;table:user_record" comment:"购买者年龄范围"`
	Gender         string `form:"gender"  search:"type:exact;column:gender;table:user_record" comment:"购买者性别"`
	CommodityId    string `form:"commodityId"  search:"type:contains;column:commodity_id;table:user_record" comment:"商品编号"`
	ActionType     string `form:"actionType"  search:"type:exact;column:action_type;table:user_record" comment:"操作行为类别"`
	UserRecordOrder
}

type UserRecordOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:user_record"`
	ItemId      string `form:"itemIdOrder"  search:"type:order;column:item_id;table:user_record"`
	ShopId      string `form:"shopIdOrder"  search:"type:order;column:shop_id;table:user_record"`
	AgeRange    string `form:"ageRangeOrder"  search:"type:order;column:age_range;table:user_record"`
	Gender      string `form:"genderOrder"  search:"type:order;column:gender;table:user_record"`
	CommodityId string `form:"commodityIdOrder"  search:"type:order;column:commodity_id;table:user_record"`
	ActionType  string `form:"actionTypeOrder"  search:"type:order;column:action_type;table:user_record"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:user_record"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:user_record"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:user_record"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:user_record"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:user_record"`
}

func (m *UserRecordGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type UserRecordInsertReq struct {
	Id          int    `json:"-" comment:"编码"` // 编码
	ItemId      string `json:"itemId" comment:"购买者编号"`
	ShopId      string `json:"shopId" comment:"商店编号"`
	AgeRange    string `json:"ageRange" comment:"购买者年龄范围"`
	Gender      string `json:"gender" comment:"购买者性别"`
	CommodityId string `json:"commodityId" comment:"商品编号"`
	ActionType  string `json:"actionType" comment:"操作行为类别"`
	common.ControlBy
}

func (s *UserRecordInsertReq) Generate(model *models.UserRecord) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ItemId = s.ItemId
	model.ShopId = s.ShopId
	model.AgeRange = s.AgeRange
	model.Gender = s.Gender
	model.CommodityId = s.CommodityId
	model.ActionType = s.ActionType
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *UserRecordInsertReq) GetId() interface{} {
	return s.Id
}

type UserRecordUpdateReq struct {
	Id          int    `uri:"id" comment:"编码"` // 编码
	ItemId      string `json:"itemId" comment:"购买者编号"`
	ShopId      string `json:"shopId" comment:"商店编号"`
	AgeRange    string `json:"ageRange" comment:"购买者年龄范围"`
	Gender      string `json:"gender" comment:"购买者性别"`
	CommodityId string `json:"commodityId" comment:"商品编号"`
	ActionType  string `json:"actionType" comment:"操作行为类别"`
	common.ControlBy
}

func (s *UserRecordUpdateReq) Generate(model *models.UserRecord) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ItemId = s.ItemId
	model.ShopId = s.ShopId
	model.AgeRange = s.AgeRange
	model.Gender = s.Gender
	model.CommodityId = s.CommodityId
	model.ActionType = s.ActionType
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *UserRecordUpdateReq) GetId() interface{} {
	return s.Id
}

// UserRecordGetReq 功能获取请求参数
type UserRecordGetReq struct {
	Id int `uri:"id"`
}

func (s *UserRecordGetReq) GetId() interface{} {
	return s.Id
}

// UserRecordDeleteReq 功能删除请求参数
type UserRecordDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *UserRecordDeleteReq) GetId() interface{} {
	return s.Ids
}
