package controller

import (
	"github.com/dsx137/anipm/anipm-server/pkg/app/entity"
	"github.com/dsx137/anipm/anipm-server/pkg/app/pojo"
	"github.com/dsx137/anipm/anipm-server/pkg/app/service"
	"github.com/dsx137/anipm/anipm-server/pkg/misc"
	"github.com/dsx137/anipm/anipm-server/pkg/util"
	"github.com/gin-gonic/gin"
)

type ControllerStaff struct {
	serv *service.ServiceStaff
}

func NewControllerStaff(g *gin.RouterGroup, serv *service.ServiceStaff) *ControllerStaff {
	ctl := &ControllerStaff{}
	ctl.serv = serv

	g.POST("/", misc.HandleController(ctl.Create))

	return ctl
}

func (ctl *ControllerStaff) Create(c *gin.Context) (*pojo.Response[string], *misc.HttpError) {
	req, err := util.ShouldBindJSON[pojo.RequestCreateProject](c)
	if err != nil {
		return nil, misc.NewHttpError(400, "Bad Request")
	}

	err = ctl.serv.Save(&entity.EntityStaff{Name: req.Name})
	if err != nil {
		return nil, misc.NewHttpError(500, "Internal Server Error")
	}

	return pojo.NewResponse("OK"), nil
}
