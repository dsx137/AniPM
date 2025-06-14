package controller

import "github.com/gin-gonic/gin"

type ControllerCut struct{}

func NewControllerCut(g *gin.RouterGroup) *ControllerCut {
	ctl := &ControllerCut{}

	return ctl
}
