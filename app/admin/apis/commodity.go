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

type Commodity struct {
	api.Api
}

// GetPage 获取商品列表
// @Summary 获取商品列表
// @Description 获取商品列表
// @Tags 商品
// @Param commodityName query string false "商品名称"
// @Param commodityBrandId query string false "商品品牌id"
// @Param commodityCategoryId query string false "商品品类id"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Commodity}} "{"code": 200, "data": [...]}"
// @Router /api/v1/commodity [get]
// @Security Bearer
func (e Commodity) GetPage(c *gin.Context) {
	req := dto.CommodityGetPageReq{}
	s := service.Commodity{}
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
	list := make([]models.Commodity, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取商品失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取商品
// @Summary 获取商品
// @Description 获取商品
// @Tags 商品
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Commodity} "{"code": 200, "data": [...]}"
// @Router /api/v1/commodity/{id} [get]
// @Security Bearer
func (e Commodity) Get(c *gin.Context) {
	req := dto.CommodityGetReq{}
	s := service.Commodity{}
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
	var object models.Commodity

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取商品失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建商品
// @Summary 创建商品
// @Description 创建商品
// @Tags 商品
// @Accept application/json
// @Product application/json
// @Param data body dto.CommodityInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/commodity [post]
// @Security Bearer
func (e Commodity) Insert(c *gin.Context) {
	req := dto.CommodityInsertReq{}
	s := service.Commodity{}
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
	if len(req.Avatar) == 0 {
		req.Avatar = "/static/uploadfile/NoImage.jpg"
	}
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建商品失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改商品
// @Summary 修改商品
// @Description 修改商品
// @Tags 商品
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.CommodityUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/commodity/{id} [put]
// @Security Bearer
func (e Commodity) Update(c *gin.Context) {
	req := dto.CommodityUpdateReq{}
	s := service.Commodity{}
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
		e.Error(500, err, fmt.Sprintf("修改商品失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除商品
// @Summary 删除商品
// @Description 删除商品
// @Tags 商品
// @Param data body dto.CommodityDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/commodity [delete]
// @Security Bearer
func (e Commodity) Delete(c *gin.Context) {
	s := service.Commodity{}
	req := dto.CommodityDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除商品失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
