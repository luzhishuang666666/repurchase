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

type RepurchaseInfo struct {
	api.Api
}

// GetPage 获取预测记录列表
// @Summary 获取预测记录列表
// @Description 获取预测记录列表
// @Tags 预测记录
// @Param modelId query string false "模型编码"
// @Param recordId query string false "用户记录编码"
// @Param status query string false "预测记录状态"
// @Param result query string false "预测结果"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RepurchaseInfo}} "{"code": 200, "data": [...]}"
// @Router /api/v1/repurchase-info [get]
// @Security Bearer
func (e RepurchaseInfo) GetPage(c *gin.Context) {
	req := dto.RepurchaseInfoGetPageReq{}
	s := service.RepurchaseInfo{}
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
	list := make([]models.RepurchaseInfo, 0)
	var count int64
	req.CreateBy = strconv.Itoa(user.GetUserId(c))
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取预测记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取预测记录
// @Summary 获取预测记录
// @Description 获取预测记录
// @Tags 预测记录
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RepurchaseInfo} "{"code": 200, "data": [...]}"
// @Router /api/v1/repurchase-info/{id} [get]
// @Security Bearer
func (e RepurchaseInfo) Get(c *gin.Context) {
	req := dto.RepurchaseInfoGetReq{}
	s := service.RepurchaseInfo{}
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
	var object models.RepurchaseInfo

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取预测记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建预测记录
// @Summary 创建预测记录
// @Description 创建预测记录
// @Tags 预测记录
// @Accept application/json
// @Product application/json
// @Param data body dto.RepurchaseInfoInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/repurchase-info [post]
// @Security Bearer
func (e RepurchaseInfo) Insert(c *gin.Context) {
	req := dto.RepurchaseInfoInsertReq{}
	s := service.RepurchaseInfo{}
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
		e.Error(500, err, fmt.Sprintf("创建预测记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改预测记录
// @Summary 修改预测记录
// @Description 修改预测记录
// @Tags 预测记录
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RepurchaseInfoUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/repurchase-info/{id} [put]
// @Security Bearer
func (e RepurchaseInfo) Update(c *gin.Context) {
	req := dto.RepurchaseInfoUpdateReq{}
	s := service.RepurchaseInfo{}
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
		e.Error(500, err, fmt.Sprintf("修改预测记录失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除预测记录
// @Summary 删除预测记录
// @Description 删除预测记录
// @Tags 预测记录
// @Param data body dto.RepurchaseInfoDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/repurchase-info [delete]
// @Security Bearer
func (e RepurchaseInfo) Delete(c *gin.Context) {
	s := service.RepurchaseInfo{}
	req := dto.RepurchaseInfoDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除预测记录失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
