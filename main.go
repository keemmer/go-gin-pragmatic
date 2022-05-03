package main

import (
	"fmt"
	"go_gin_pragmatic/controller"
	"go_gin_pragmatic/middleware"
	"go_gin_pragmatic/repository"
	"go_gin_pragmatic/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/tpkeeper/gin-dump"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService service.VideoService = service.New(videoRepository)
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func main() {
	defer videoRepository.CloseDB()
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

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("api", middleware.AuthorizeJWT(), middleware.Admin())
	{
		apiRoutes.GET("/video", func(c *gin.Context) {
			fmt.Println(c)
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
		viewRoutes.GET("/videos", videoController.ShowView)
	}

	server.Run(":8000")
}
