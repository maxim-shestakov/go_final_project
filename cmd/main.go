package main

import (
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	gin "github.com/gin-gonic/gin"
	"github.com/maxim-shestakov/final-yandex-project/internal/db"
	"github.com/maxim-shestakov/final-yandex-project/pkg/handlers"
	"github.com/maxim-shestakov/final-yandex-project/pkg/repo"
	"github.com/maxim-shestakov/final-yandex-project/pkg/service"

)

func main() {
	dbsqlite, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbsqlite.Close()

	Repository := repo.New(dbsqlite)
	Service := service.New(Repository)
	Handlers := handlers.New(Service)

	appPort := ":" + os.Getenv("TODO_PORT")
	if appPort == ":" {
		appPort = ":7540"
	}

	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	Handlers.InitRoutes(r)
	r.GET("/", Handlers.Index)
	r.GET("api/nextdate", Handlers.NextDate)
	r.POST("/api/task", Handlers.CreateTask)
	r.GET("/api/tasks", Handlers.GetTasks)
	r.GET("/api/task", Handlers.GetTaskById)
	r.PUT("/api/task", Handlers.UpdateTask)
	r.POST("/api/task/done", Handlers.DoneTask)
	r.DELETE("/api/task", Handlers.DeleteTask)

	err = r.Run(appPort)
	if err != nil {
		panic(err)
	}
}
