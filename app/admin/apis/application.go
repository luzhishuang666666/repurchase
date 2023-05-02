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

type Application struct {
	api.Api
}

// GetPage 获取申请列表
// @Summary 获取申请列表
// @Description 获取申请列表
// @Tags 申请
// @Param type query string false "申请类型"
// @Param applicationContext query string false "申请内容"
// @Param status query string false "状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Application}} "{"code": 200, "data": [...]}"
// @Router /api/v1/application [get]
// @Security Bearer
func (e Application) GetPage(c *gin.Context) {
	req := dto.ApplicationGetPageReq{}
	s := service.Application{}
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
	list := make([]models.Application, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取申请失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取申请
// @Summary 获取申请
// @Description 获取申请
// @Tags 申请
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Application} "{"code": 200, "data": [...]}"
// @Router /api/v1/application/{id} [get]
// @Security Bearer
func (e Application) Get(c *gin.Context) {
	req := dto.ApplicationGetReq{}
	s := service.Application{}
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
	var object models.Application

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取申请失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建申请
// @Summary 创建申请
// @Description 创建申请
// @Tags 申请
// @Accept application/json
// @Product application/json
// @Param data body dto.ApplicationInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/application [post]
// @Security Bearer
func (e Application) Insert(c *gin.Context) {
	req := dto.ApplicationInsertReq{}
	s := service.Application{}
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
		e.Error(500, err, fmt.Sprintf("创建申请失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改申请
// @Summary 修改申请
// @Description 修改申请
// @Tags 申请
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ApplicationUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/application/{id} [put]
// @Security Bearer
func (e Application) Update(c *gin.Context) {
	req := dto.ApplicationUpdateReq{}
	s := service.Application{}
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
		e.Error(500, err, fmt.Sprintf("修改申请失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除申请
// @Summary 删除申请
// @Description 删除申请
// @Tags 申请
// @Param data body dto.ApplicationDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/application [delete]
// @Security Bearer
func (e Application) Delete(c *gin.Context) {
	s := service.Application{}
	req := dto.ApplicationDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除申请失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// Approval 审批申请
// @Summary 审批申请
// @Description 审批申请
// @Tags 申请
// @Router /api/v1/application/approval [GET]
// @Security Bearer
func (e Application) Approval(c *gin.Context) {
	req := dto.ApplicationApprovalReq{}
	s := service.Application{}
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

	err = s.Approval(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建申请失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}
