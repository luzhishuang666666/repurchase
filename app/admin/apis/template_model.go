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

type TemplateModel struct {
	api.Api
}

// GetPage 获取模型模板列表
// @Summary 获取模型模板列表
// @Description 获取模型模板列表
// @Tags 模型模板
// @Param templateName query string false "模板名称"
// @Param templateType query string false "模板类型"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.TemplateModel}} "{"code": 200, "data": [...]}"
// @Router /api/v1/template-model [get]
// @Security Bearer
func (e TemplateModel) GetPage(c *gin.Context) {
	req := dto.TemplateModelGetPageReq{}
	s := service.TemplateModel{}
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
	list := make([]models.TemplateModel, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取模型模板失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取模型模板
// @Summary 获取模型模板
// @Description 获取模型模板
// @Tags 模型模板
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.TemplateModel} "{"code": 200, "data": [...]}"
// @Router /api/v1/template-model/{id} [get]
// @Security Bearer
func (e TemplateModel) Get(c *gin.Context) {
	req := dto.TemplateModelGetReq{}
	s := service.TemplateModel{}
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
	var object models.TemplateModel

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取模型模板失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建模型模板
// @Summary 创建模型模板
// @Description 创建模型模板
// @Tags 模型模板
// @Accept application/json
// @Product application/json
// @Param data body dto.TemplateModelInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/template-model [post]
// @Security Bearer
func (e TemplateModel) Insert(c *gin.Context) {
	req := dto.TemplateModelInsertReq{}
	s := service.TemplateModel{}
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
		e.Error(500, err, fmt.Sprintf("创建模型模板失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改模型模板
// @Summary 修改模型模板
// @Description 修改模型模板
// @Tags 模型模板
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.TemplateModelUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/template-model/{id} [put]
// @Security Bearer
func (e TemplateModel) Update(c *gin.Context) {
	req := dto.TemplateModelUpdateReq{}
	s := service.TemplateModel{}
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
		e.Error(500, err, fmt.Sprintf("修改模型模板失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除模型模板
// @Summary 删除模型模板
// @Description 删除模型模板
// @Tags 模型模板
// @Param data body dto.TemplateModelDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/template-model [delete]
// @Security Bearer
func (e TemplateModel) Delete(c *gin.Context) {
	s := service.TemplateModel{}
	req := dto.TemplateModelDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除模型模板失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
