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
)

type ShopCommodity struct {
	api.Api
}

// GetPage 获取商店商品汇总列表
// @Summary 获取商店商品汇总列表
// @Description 获取商店商品汇总列表
// @Tags 商店商品汇总
// @Param commodityId query string false "商品编号"
// @Param shopId query string false "商店编号"
// @Param status query string false "状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ShopCommodity}} "{"code": 200, "data": [...]}"
// @Router /api/v1/shop-commodity [get]
// @Security Bearer
func (e ShopCommodity) GetPage(c *gin.Context) {
	req := dto.ShopCommodityGetPageReq{}
	s := service.ShopCommodity{}
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
	list := make([]models.ShopCommodity, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取商店商品汇总失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取商店商品汇总
// @Summary 获取商店商品汇总
// @Description 获取商店商品汇总
// @Tags 商店商品汇总
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.ShopCommodity} "{"code": 200, "data": [...]}"
// @Router /api/v1/shop-commodity/{id} [get]
// @Security Bearer
func (e ShopCommodity) Get(c *gin.Context) {
	req := dto.ShopCommodityGetReq{}
	s := service.ShopCommodity{}
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
	var object models.ShopCommodity

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取商店商品汇总失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建商店商品汇总
// @Summary 创建商店商品汇总
// @Description 创建商店商品汇总
// @Tags 商店商品汇总
// @Accept application/json
// @Product application/json
// @Param data body dto.ShopCommodityInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/shop-commodity [post]
// @Security Bearer
func (e ShopCommodity) Insert(c *gin.Context) {
	req := dto.ShopCommodityInsertReq{}
	s := service.ShopCommodity{}
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

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建商店商品汇总失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改商店商品汇总
// @Summary 修改商店商品汇总
// @Description 修改商店商品汇总
// @Tags 商店商品汇总
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ShopCommodityUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/shop-commodity/{id} [put]
// @Security Bearer
func (e ShopCommodity) Update(c *gin.Context) {
	req := dto.ShopCommodityUpdateReq{}
	s := service.ShopCommodity{}
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
		e.Error(500, err, fmt.Sprintf("修改商店商品汇总失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除商店商品汇总
// @Summary 删除商店商品汇总
// @Description 删除商店商品汇总
// @Tags 商店商品汇总
// @Param data body dto.ShopCommodityDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/shop-commodity [delete]
// @Security Bearer
func (e ShopCommodity) Delete(c *gin.Context) {
	s := service.ShopCommodity{}
	req := dto.ShopCommodityDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除商店商品汇总失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
