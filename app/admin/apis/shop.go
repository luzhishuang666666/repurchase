package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	"go-admin/common/middleware"
	"strconv"
)

type Shop struct {
	api.Api
}

// GetPage 获取商店列表
// @Summary 获取商店列表
// @Description 获取商店列表
// @Tags 商店
// @Param shopNo query string false "商店编号"
// @Param shopName query string false "商店名称"
// @Param shopRemark query string false "商店备注"
// @Param status query string false "状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Shop}} "{"code": 200, "data": [...]}"
// @Router /api/v1/shop [get]
// @Security Bearer
func (e Shop) GetPage(c *gin.Context) {
	req := dto.ShopGetPageReq{}
	s := service.Shop{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.Shop, 0)
	var count int64
	req.CreateBy = strconv.Itoa(user.GetUserId(c))

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取商店失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取商店
// @Summary 获取商店
// @Description 获取商店
// @Tags 商店
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Shop} "{"code": 200, "data": [...]}"
// @Router /api/v1/shop/{id} [get]
// @Security Bearer
func (e Shop) Get(c *gin.Context) {
	req := dto.ShopGetReq{}
	s := service.Shop{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Shop

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取商店失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建商店
// @Summary 创建商店
// @Description 创建商店
// @Tags 商店
// @Accept application/json
// @Product application/json
// @Param data body dto.ShopInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/shop [post]
// @Security Bearer
func (e Shop) Insert(c *gin.Context) {
	req := dto.ShopInsertReq{}
	s := service.Shop{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	// 生成对应的编号
	shopNo, err := middleware.GenerateDistributedId()
	if err != nil {
		e.Error(500, err, "商店编号生成失败")
		return
	}
	req.SetShopNo(shopNo)

	// 设置状态
	req.SetStatus("2")

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建商店失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改商店
// @Summary 修改商店
// @Description 修改商店
// @Tags 商店
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ShopUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/shop/{id} [put]
// @Security Bearer
func (e Shop) Update(c *gin.Context) {
	req := dto.ShopUpdateReq{}
	s := service.Shop{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改商店失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除商店
// @Summary 删除商店
// @Description 删除商店
// @Tags 商店
// @Param data body dto.ShopDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/shop [delete]
// @Security Bearer
func (e Shop) Delete(c *gin.Context) {
	s := service.Shop{}
	req := dto.ShopDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除商店失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

func (e Shop) ChangeStatus(c *gin.Context) {
	req := dto.ShopChangeStatusReq{}
	s := service.Shop{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.ChangeStatus(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("商店状态变更失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "查询成功")
}

func (e Shop) ShopRank(c *gin.Context) {
	req := dto.ShopRankReq{}
	s := service.Shop{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	shopRank := dto.ShopRankResp{}

	err = s.ShopRank(&req, &shopRank)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("商店排行榜信息查询失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(shopRank, "查询成功")
}

func (e Shop) ShopAnalise(c *gin.Context) {

	req := dto.ShopAnaliseReq{}
	s := service.Shop{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	shopAnalise := dto.ShopAnaliseResp{}

	err = s.ShopAnalise(&req, &shopAnalise)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("商店排行榜信息查询失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(shopAnalise, "查询成功")
}
