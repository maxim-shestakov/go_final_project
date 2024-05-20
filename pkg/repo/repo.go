package repo

import (
	"database/sql"

	"github.com/maxim-shestakov/final-yandex-project/internal/models"
)

type RepoInterface interface {
	GetTasks() ([]models.Task, error)
	CreateTask(task *models.Task) (int64, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id int64) error
	GetTaskById(id int64) (models.Task, error)
	
}

type Repo struct {
	RepoInterface
}

func New(db *sql.DB) *Repo {
	return &Repo{
		RepoInterface: NewSqliteRepo(db),
	}
}
