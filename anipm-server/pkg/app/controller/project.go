package controller

import (
	"fmt"

	"github.com/dsx137/anipm/anipm-server/pkg/app/entity"
	"github.com/dsx137/anipm/anipm-server/pkg/app/pojo"
	"github.com/dsx137/anipm/anipm-server/pkg/app/service"
	"github.com/dsx137/anipm/anipm-server/pkg/misc"
	"github.com/dsx137/anipm/anipm-server/pkg/util"
	"github.com/gin-gonic/gin"
)

type ControllerProject struct {
	serv *service.ServiceProject
}

func NewControllerProject(g *gin.RouterGroup, serv *service.ServiceProject) *ControllerProject {
	ctl := &ControllerProject{}
	ctl.serv = serv

	g.GET("/", misc.HandleController(ctl.List))
	g.POST("/", misc.HandleController(ctl.Create))

	return ctl
}

func (ctl *ControllerProject) Create(c *gin.Context) (*pojo.Response[string], *misc.HttpError) {
	req, err := util.ShouldBindJSON[pojo.RequestCreateProject](c)
	if err != nil {
		return nil, misc.NewHttpError(400, "Bad Request")
	}

	if ctl.serv.ExistByName(req.Name) {
		return nil, misc.NewHttpError(400, "项目已存在")
	}
	err = ctl.serv.Save(&entity.EntityProject{Name: req.Name})
	if err != nil {
		return nil, misc.NewHttpError(500, "Internal Server Error")
	}

	return pojo.OK, nil
}

func (ctl *ControllerProject) List(c *gin.Context) (*pojo.Response[[]pojo.RequestCreateProject], *misc.HttpError) {
	projects := []pojo.RequestCreateProject{}

	ps, err := ctl.serv.FindAll()
	if err != nil {
		return nil, misc.NewHttpError(500, fmt.Sprintf("获取项目列表失败: %s", err))
	}

	for _, p := range ps {
		projects = append(projects, pojo.RequestCreateProject{Name: p.Name})
	}

	return pojo.NewResponse(projects), nil
}
