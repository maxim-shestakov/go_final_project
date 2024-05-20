package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maxim-shestakov/final-yandex-project/pkg/service"
)

type HandlersInterface interface {
	GetTasks(c *gin.Context)
	GetTaskById(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	DoneTask(c *gin.Context)
	InitRoutes(mux *gin.Engine)
	Index(c *gin.Context)
	NextDate(c *gin.Context)
}

type Handlers struct {
	HandlersInterface
}

func New(services *service.Service) *Handlers {
	return &Handlers{
		HandlersInterface: NewController(services),
	}
}
