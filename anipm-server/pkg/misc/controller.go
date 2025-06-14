package misc

import (
	"net/http"

	"github.com/dsx137/anipm/anipm-server/pkg/app/pojo"
	"github.com/gin-gonic/gin"
)

type Nothing any
type Controller[T any] func(c *gin.Context) (*pojo.Response[T], *HttpError)

func HandleController[T any](controller Controller[T]) func(c *gin.Context) {
	return func(c *gin.Context) {
		response, err := controller(c)
		if err != nil {
			c.JSON(err.StatusCode, pojo.NewResponse(err.Error()))
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
