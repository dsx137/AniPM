package controller

import "github.com/gin-gonic/gin"

type ControllerStep struct{}

func NewControllerStep(g *gin.RouterGroup) *ControllerStep {
	ctl := &ControllerStep{}

	return ctl
}
