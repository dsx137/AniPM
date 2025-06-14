package util

import "github.com/gin-gonic/gin"

func ShouldBindJSON[T any](c *gin.Context) (*T, error) {
	req := new(T)
	return req, c.ShouldBindJSON(req)
}
