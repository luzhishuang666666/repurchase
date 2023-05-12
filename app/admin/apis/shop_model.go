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

type ShopModel struct {
	api.Api
}

// GetPage 获取复购预测模型列表
// @Summary 获取复购预测模型列表
// @Description 获取复购预测模型列表
// @Tags 复购预测模型
// @Param shopId query string false "商店编码"
// @Param templateId query string false "模板编码"
// @Param modelName query string false "模型名称"
// @Param modelRemark query string false "模型备注"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ShopModel}} "{"code": 200, "data": [...]}"
// @Router /api/v1/shop-model [get]
// @Security Bearer
func (e ShopModel) GetPage(c *gin.Context) {
	req := dto.ShopModelGetPageReq{}
	s := service.ShopModel{}
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
	list := make([]models.ShopModel, 0)
	var count int64
	req.CreateBy = strconv.Itoa(user.GetUserId(c))
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取复购预测模型失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取复购预测模型
// @Summary 获取复购预测模型
// @Description 获取复购预测模型
// @Tags 复购预测模型
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.ShopModel} "{"code": 200, "data": [...]}"
// @Router /api/v1/shop-model/{id} [get]
// @Security Bearer
func (e ShopModel) Get(c *gin.Context) {
	req := dto.ShopModelGetReq{}
	s := service.ShopModel{}
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
	var object models.ShopModel

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取复购预测模型失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建复购预测模型
// @Summary 创建复购预测模型
// @Description 创建复购预测模型
// @Tags 复购预测模型
// @Accept application/json
// @Product application/json
// @Param data body dto.ShopModelInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/shop-model [post]
// @Security Bearer
func (e ShopModel) Insert(c *gin.Context) {
	req := dto.ShopModelInsertReq{}
	s := service.ShopModel{}
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
		e.Error(500, err, fmt.Sprintf("创建复购预测模型失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改复购预测模型
// @Summary 修改复购预测模型
// @Description 修改复购预测模型
// @Tags 复购预测模型
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ShopModelUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/shop-model/{id} [put]
// @Security Bearer
func (e ShopModel) Update(c *gin.Context) {
	req := dto.ShopModelUpdateReq{}
	s := service.ShopModel{}
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
		e.Error(500, err, fmt.Sprintf("修改复购预测模型失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除复购预测模型
// @Summary 删除复购预测模型
// @Description 删除复购预测模型
// @Tags 复购预测模型
// @Param data body dto.ShopModelDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/shop-model [delete]
// @Security Bearer
func (e ShopModel) Delete(c *gin.Context) {
	s := service.ShopModel{}
	req := dto.ShopModelDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除复购预测模型失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

func (e ShopModel) Forecast(c *gin.Context) {
	s := service.ShopModel{}
	req := dto.ShopModelForecastReq{}
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

	err = s.Forecast(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("复购预测失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK("复购预测成功", "复购预测成功")
}
