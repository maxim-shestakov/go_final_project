package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maxim-shestakov/final-yandex-project/internal/models"
	"github.com/maxim-shestakov/final-yandex-project/internal/repeat"
	"github.com/maxim-shestakov/final-yandex-project/pkg/service"
)

type Controller struct {
	service service.ServiceInterface
}

func NewController(service service.ServiceInterface) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) GetTasks(c *gin.Context) {
	tasks, err := ctrl.service.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (ctrl *Controller) CreateTask(c *gin.Context) {
	var task models.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var id int64
	if id, err = ctrl.service.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (ctrl *Controller) UpdateTask(c *gin.Context) {
	var task models.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ctrl.service.UpdateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (ctrl *Controller) DeleteTask(c *gin.Context) {
	ids := c.Query("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.service.DeleteTask(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (ctrl *Controller) GetTaskById(c *gin.Context) {
	ids := c.Query("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := ctrl.service.GetTaskById(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (ctrl *Controller) DoneTask(c *gin.Context) {
	ids := c.Query("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := ctrl.service.GetTaskById(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if task.Repeat != "" {
		task.Date, err = repeat.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = ctrl.service.UpdateTask(&task)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	err = ctrl.service.DeleteTask(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (ctrl *Controller) InitRoutes(mux *gin.Engine) {
	mux.StaticFS("/web", gin.Dir("./web", true))
	mux.StaticFile("/favicon.ico", "./web/favicon.ico")
	mux.StaticFile("/index.html", "./web/index.html")
	mux.Static("/css", "./web/css")
	mux.Static("/js", "./web/js")
	mux.StaticFile("/login.html", "./web/login.html")
}

func (ctrl *Controller) Index(c *gin.Context) {
	c.File("./web/index.html")
}

func (ctrl *Controller) NextDate(c *gin.Context) {
	now := c.Query("now")
	date := c.Query("date")
	rpt := c.Query("repeat")
	nowTime, err := time.Parse(`20060102`, now)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	nxtDate, err := repeat.NextDate(nowTime, date, rpt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, nxtDate)

}
