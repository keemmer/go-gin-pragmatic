package controller

import (
	"go_gin_pragmatic/entity"
	"go_gin_pragmatic/service"
	"go_gin_pragmatic/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

type VideoController interface {
	FindAll() []entity.Video
	// Save(ctx *gin.Context) entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool",validators.ValidateCoolTitle)
	return &controller{service: service}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

// func (c *controller) Save(ctx *gin.Context) entity.Video {
func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	// ctx.BindJSON(&video)
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context){
	videos:=c.service.FindAll()
	data:= gin.H{
		"title":"Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html",data)
}

