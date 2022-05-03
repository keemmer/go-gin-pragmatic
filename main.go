package main

import (
	"go_gin_pragmatic/controller"
	"go_gin_pragmatic/middleware"
	"go_gin_pragmatic/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func main() {
	setupLogOutput()

	// server := gin.Default()
	server := gin.New()
	// server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth(), gindump.Dump())
	server.Use(gin.Recovery(), middleware.Logger())

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello GIN",
		})
	})

	server.Static("/css", "./templates/css/")
	server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("api")
	{
		apiRoutes.GET("/video", func(c *gin.Context) {
			c.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/video", func(c *gin.Context) {
			err := videoController.Save(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {

				c.JSON(http.StatusOK, gin.H{"message": "Video input is Valid!!"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos",videoController.ShowAll)
	}

	server.Run(":8000")
}
