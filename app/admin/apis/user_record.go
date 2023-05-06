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

type UserRecord struct {
	api.Api
}

// GetPage 获取用户记录列表
// @Summary 获取用户记录列表
// @Description 获取用户记录列表
// @Tags 用户记录
// @Param itemId query string false "购买者编号"
// @Param shopId query string false "商店编号"
// @Param ageRange query string false "购买者年龄范围"
// @Param gender query string false "购买者性别"
// @Param commodityId query string false "商品编号"
// @Param actionType query string false "操作行为类别"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.UserRecord}} "{"code": 200, "data": [...]}"
// @Router /api/v1/user-record [get]
// @Security Bearer
func (e UserRecord) GetPage(c *gin.Context) {
	req := dto.UserRecordGetPageReq{}
	s := service.UserRecord{}
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
	list := make([]models.UserRecord, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取用户记录
// @Summary 获取用户记录
// @Description 获取用户记录
// @Tags 用户记录
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.UserRecord} "{"code": 200, "data": [...]}"
// @Router /api/v1/user-record/{id} [get]
// @Security Bearer
func (e UserRecord) Get(c *gin.Context) {
	req := dto.UserRecordGetReq{}
	s := service.UserRecord{}
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
	var object models.UserRecord

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建用户记录
// @Summary 创建用户记录
// @Description 创建用户记录
// @Tags 用户记录
// @Accept application/json
// @Product application/json
// @Param data body dto.UserRecordInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/user-record [post]
// @Security Bearer
func (e UserRecord) Insert(c *gin.Context) {
	req := dto.UserRecordInsertReq{}
	s := service.UserRecord{}
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
		e.Error(500, err, fmt.Sprintf("创建用户记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改用户记录
// @Summary 修改用户记录
// @Description 修改用户记录
// @Tags 用户记录
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.UserRecordUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/user-record/{id} [put]
// @Security Bearer
func (e UserRecord) Update(c *gin.Context) {
	req := dto.UserRecordUpdateReq{}
	s := service.UserRecord{}
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
		e.Error(500, err, fmt.Sprintf("修改用户记录失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除用户记录
// @Summary 删除用户记录
// @Description 删除用户记录
// @Tags 用户记录
// @Param data body dto.UserRecordDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/user-record [delete]
// @Security Bearer
func (e UserRecord) Delete(c *gin.Context) {
	s := service.UserRecord{}
	req := dto.UserRecordDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除用户记录失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
