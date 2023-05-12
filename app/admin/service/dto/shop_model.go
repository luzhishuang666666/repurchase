package dto

import (
	_ "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ShopModelGetPageReq struct {
	dto.Pagination `search:"-"`
	ShopId         string `form:"shopId"  search:"type:exact;column:shop_id;table:shop_model" comment:"商店编码"`
	TemplateId     string `form:"templateId"  search:"type:exact;column:template_id;table:shop_model" comment:"模板编码"`
	ModelName      string `form:"modelName"  search:"type:contains;column:model_name;table:shop_model" comment:"模型名称"`
	ModelRemark    string `form:"modelRemark"  search:"type:contains;column:model_remark;table:shop_model" comment:"模型备注"`
	ShopModelOrder
}

type ShopModelOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:shop_model"`
	ShopId      string `form:"shopIdOrder"  search:"type:order;column:shop_id;table:shop_model"`
	TemplateId  string `form:"templateIdOrder"  search:"type:order;column:template_id;table:shop_model"`
	ModelName   string `form:"modelNameOrder"  search:"type:order;column:model_name;table:shop_model"`
	ModelRemark string `form:"modelRemarkOrder"  search:"type:order;column:model_remark;table:shop_model"`
	ModelParam  string `form:"modelParamOrder"  search:"type:order;column:model_param;table:shop_model"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:shop_model"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:shop_model"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:shop_model"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:shop_model"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:shop_model"`
}

func (m *ShopModelGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ShopModelInsertReq struct {
	Id          int    `json:"-" comment:"编码"` // 编码
	ShopId      string `json:"shopId" comment:"商店编码"`
	TemplateId  string `json:"templateId" comment:"模板编码"`
	ModelName   string `json:"modelName" comment:"模型名称"`
	ModelRemark string `json:"modelRemark" comment:"模型备注"`
	ModelParam  string `json:"modelParam" comment:"模型参数"`
	common.ControlBy
}

func (s *ShopModelInsertReq) Generate(model *models.ShopModel) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ShopId = s.ShopId
	model.TemplateId = s.TemplateId
	model.ModelName = s.ModelName
	model.ModelRemark = s.ModelRemark
	model.ModelParam = s.ModelParam
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *ShopModelInsertReq) GetId() interface{} {
	return s.Id
}

type ShopModelUpdateReq struct {
	Id          int    `uri:"id" comment:"编码"` // 编码
	ShopId      string `json:"shopId" comment:"商店编码"`
	TemplateId  string `json:"templateId" comment:"模板编码"`
	ModelName   string `json:"modelName" comment:"模型名称"`
	ModelRemark string `json:"modelRemark" comment:"模型备注"`
	ModelParam  string `json:"modelParam" comment:"模型参数"`
	common.ControlBy
}

func (s *ShopModelUpdateReq) Generate(model *models.ShopModel) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ShopId = s.ShopId
	model.TemplateId = s.TemplateId
	model.ModelName = s.ModelName
	model.ModelRemark = s.ModelRemark
	model.ModelParam = s.ModelParam
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *ShopModelUpdateReq) GetId() interface{} {
	return s.Id
}

// ShopModelGetReq 功能获取请求参数
type ShopModelGetReq struct {
	Id int `uri:"id"`
}

func (s *ShopModelGetReq) GetId() interface{} {
	return s.Id
}

// ShopModelDeleteReq 功能删除请求参数
type ShopModelDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ShopModelDeleteReq) GetId() interface{} {
	return s.Ids
}

type ShopModelForecastReq struct {
	Id           int    `json:"id" comment:"编码"` // 编码
	UserRecordId string `json:"userRecordId" comment:"用户记录编码"`
	ModelName    string `json:"modelName" comment:"模型名称"`
}

type ShopForecastMq struct {
	ModelId          int              `json:"model_id"`
	RepurchaseInfoId int              `json:"repurchase_info_id"`
	ModelType        string           `json:"model_type"`
	ModelParam       string           `json:"model_param"`
	ModelData        ShopForecastData `json:"model_data"`
}

type ShopForecastData struct {
	UserId                string `json:"user_id"`
	MerchantId            string `json:"merchant_id"`
	AgeRange              string `json:"age_range"`
	Gender                string `json:"gender"`
	ClickInEleven         int    `json:"click_in_eleven"`
	ClickNotInEleven      int    `json:"clickNotInEleven"`
	BuyInEleven           int    `json:"buyInEleven"`
	BuyNotInEleven        int    `json:"buyNotInEleven"`
	CollectionInEleven    int    `json:"collectionInEleven"`
	CollectionNotInEleven int    `json:"collectionNotInEleven"`
	InElevenLogNums       int    `json:"inElevenLogNums"`
	AllLogNums            int    `json:"allLogNums"`
}
