package apis

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type Brand struct {
	api.Api
}

// GetPage 获取品牌列表
// @Summary 获取品牌列表
// @Description 获取品牌列表
// @Tags 品牌
// @Param brandName query string false "品牌名称"
// @Param brandRemark query string false "品牌备注"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Brand}} "{"code": 200, "data": [...]}"
// @Router /api/v1/brand [get]
// @Security Bearer
func (e Brand) GetPage(c *gin.Context) {
	req := dto.BrandGetPageReq{}
	s := service.Brand{}
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
	list := make([]models.Brand, 0)
	var count int64
	req.CreateBy = strconv.Itoa(user.GetUserId(c))
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取品牌失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取品牌
// @Summary 获取品牌
// @Description 获取品牌
// @Tags 品牌
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Brand} "{"code": 200, "data": [...]}"
// @Router /api/v1/brand/{id} [get]
// @Security Bearer
func (e Brand) Get(c *gin.Context) {
	req := dto.BrandGetReq{}
	s := service.Brand{}
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
	var object models.Brand

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取品牌失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建品牌
// @Summary 创建品牌
// @Description 创建品牌
// @Tags 品牌
// @Accept application/json
// @Product application/json
// @Param data body dto.BrandInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/brand [post]
// @Security Bearer
func (e Brand) Insert(c *gin.Context) {
	req := dto.BrandInsertReq{}
	s := service.Brand{}
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
		e.Error(500, err, fmt.Sprintf("创建品牌失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改品牌
// @Summary 修改品牌
// @Description 修改品牌
// @Tags 品牌
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.BrandUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/brand/{id} [put]
// @Security Bearer
func (e Brand) Update(c *gin.Context) {
	req := dto.BrandUpdateReq{}
	s := service.Brand{}
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
		e.Error(500, err, fmt.Sprintf("修改品牌失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除品牌
// @Summary 删除品牌
// @Description 删除品牌
// @Tags 品牌
// @Param data body dto.BrandDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/brand [delete]
// @Security Bearer
func (e Brand) Delete(c *gin.Context) {
	s := service.Brand{}
	req := dto.BrandDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除品牌失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
