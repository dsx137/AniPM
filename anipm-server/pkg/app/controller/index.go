package controller

import (
	"github.com/dsx137/anipm/anipm-server/pkg/app/pojo"
	"github.com/dsx137/anipm/anipm-server/pkg/misc"
	"github.com/gin-gonic/gin"
)

type ControllerIndex struct{}

func NewControllerIndex(g *gin.RouterGroup) *ControllerIndex {
	ctl := &ControllerIndex{}

	g.GET("/", misc.HandleController(ctl.Index))

	return ctl
}

func (ctl *ControllerIndex) Index(c *gin.Context) (*pojo.Response[string], *misc.HttpError) {
	return pojo.NewResponse("Hello, AniOA!"), nil
}
