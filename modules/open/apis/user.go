package apis

import (
	"github.com/mooncake9527/orange/modules/open/models"
	"github.com/mooncake9527/orange/modules/open/service"
	"github.com/mooncake9527/orange/modules/open/service/dto"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/mooncake9527/orange-core/common/apis"
	"github.com/mooncake9527/orange-core/core/base"
)

type UserApi struct {
	apis.NApi
}

var ApiUser = UserApi{}

// QueryPage 获取用户表列表
// @Summary 获取用户表列表
// @Tags open-User
// @Accept application/json
// @Product application/json
// @Param data body dto.UserGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.User}} "{"code": 200, "data": [...]}"
// @Router /orange/v1/open/user/page [post]
// @Security Bearer
func (e *UserApi) QueryPage(c *gin.Context) {
	var req dto.UserGetPageReq
	if err := e.Bind(c, &req); err != nil {
		e.ParamError(c, err)
		return
	}
	if err := req.Valid(); err != nil {
		e.ParamError(c, err)
		return
	}
	list := make([]models.User, 0, req.GetSize())
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerUser.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.PageOK(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取用户表
// @Summary 获取用户表
// @Tags open-User
// @Accept application/json
// @Product application/json
// @Param data body dto.UserReq true "body"
// @Success 200 {object} base.Resp{data=models.User} "{"code": 200, "data": [...]}"
// @Router /orange/v1/open/user/get [post]
// @Security Bearer
func (e *UserApi) Get(c *gin.Context) {
	var req dto.UserReq
	if err := e.Bind(c, &req); err != nil {
		e.ParamError(c, err)
		return
	}
	if err := req.Valid(); err != nil {
		e.ParamError(c, err)
		return
	}
	var condition models.User
	_ = copier.CopyWithOption(&condition, &req, copier.Option{IgnoreEmpty: true})
	var data models.User
	if err := service.SerUser.Get(&condition, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.OK(c, data)
}

// Create 创建用户表
// @Summary 创建用户表
// @Tags open-User
// @Accept application/json
// @Product application/json
// @Param data body dto.UserReq true "body"
// @Success 200 {object} base.Resp{data=models.User} "{"code": 200, "data": [...]}"
// @Router /orange/v1/open/user/create [post]
// @Security Bearer
func (e *UserApi) Create(c *gin.Context) {
	var req dto.UserReq
	if err := e.Bind(c, &req); err != nil {
		e.ParamError(c, err)
		return
	}
	if err := req.Valid(); err != nil {
		e.ParamError(c, err)
		return
	}
	var data models.User
	_ = copier.Copy(&data, req)
	if err := service.SerUser.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.OK(c, data)
}

// Update 更新用户表
// @Summary 更新用户表
// @Tags open-User
// @Accept application/json
// @Product application/json
// @Success 200 {object} base.Resp{data=models.User} "{"code": 200, "data": [...]}"
// @Router /orange/v1/open/user/update [post]
// @Security Bearer
func (e *UserApi) Update(c *gin.Context) {
	var req dto.UserReq
	if err := e.Bind(c, &req); err != nil {
		e.ParamError(c, err)
		return
	}
	if err := req.Valid(); err != nil {
		e.ParamError(c, err)
		return
	}
	var data models.User
	_ = copier.Copy(&data, req)
	if err := service.SerUser.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.OK(c, data)
}

// Del 删除用户表
// @Summary 删除用户表
// @Tags open-User
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /orange/v1/open/user/del [post]
// @Security Bearer
func (e *UserApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := e.Bind(c, &req); err != nil {
		e.ParamError(c, err)
		return
	}
	if err := req.Valid(); err != nil {
		e.ParamError(c, err)
		return
	}
	if err := service.SerUser.DelIds(&models.User{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.OK(c, nil)
}
