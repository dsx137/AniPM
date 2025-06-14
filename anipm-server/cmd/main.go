package cmd

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/dsx137/anipm/anipm-server/pkg/app/config"
	"github.com/dsx137/anipm/anipm-server/pkg/app/controller"
	"github.com/dsx137/anipm/anipm-server/pkg/app/pojo"
	"github.com/dsx137/anipm/anipm-server/pkg/app/repository"
	"github.com/dsx137/anipm/anipm-server/pkg/app/service"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine, clientFS embed.FS) {
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, pojo.NewResponse("Method Not Allowed"))
	})

	clientDistFS, _ := fs.Sub(clientFS, "anipm-client/dist")
	r.GET("/assets/*filepath", gin.WrapH(http.FileServer(http.FS(clientDistFS))))

	r.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, "/api") {
			c.Status(404)
			return
		}
		indexFile, err := clientDistFS.Open("index.html")
		if err != nil {
			c.Status(500)
			return
		}
		defer indexFile.Close()
		io.Copy(c.Writer, indexFile)
	})

	api := r.Group("/api")
	{
		controller.NewControllerIndex(api.Group("/"))
		controller.NewControllerProject(api.Group("/projects"), service.NewServiceProject(repository.NewRepositoryProject(config.BaseDir)))
		controller.NewControllerCut(api.Group("/projects/:projectId/cuts"))
		controller.NewControllerStep(api.Group("/projects/:projectId/cuts/:cutId/steps"))
		controller.NewControllerStaff(api.Group("/staffs"), service.NewServiceStaff(repository.NewRepositoryStaff(config.BaseDir)))
	}
}

func Main(clientFS embed.FS) {
	if mode := os.Getenv("GIN_MODE"); mode == "" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(mode)
	}

	r := gin.Default()

	initRouter(r, clientFS)

	r.Run(":8080")
}
