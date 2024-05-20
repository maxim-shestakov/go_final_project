package service

import (
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/maxim-shestakov/final-yandex-project/internal/models"
	"github.com/maxim-shestakov/final-yandex-project/internal/repeat"
	"github.com/maxim-shestakov/final-yandex-project/pkg/repo"
)

type todoService struct {
	repo repo.RepoInterface
}

func NewService(r repo.RepoInterface) *todoService {
	return &todoService{
		repo: r,
	}
}

func (s *todoService) GetTasks() ([]models.Task, error) {
	tasks, err := s.repo.GetTasks()
	if err != nil {
		return nil, errors.WithMessage(err, "service: get tasks")
	}
	if len(tasks) == 0 {
		return []models.Task{}, nil
	}
	return tasks, nil
}

func (s *todoService) CreateTask(task *models.Task) (int64, error) {
	if task == nil {
		return 0, errors.New("task is nil")
	}

	if task.Title == "" {
		return 0, errors.New("task title is empty")
	}

	if task.Date == "" {
		task.Date = time.Now().Format("20060102")
	}
	_, dateCheck := time.Parse("20060102", task.Date)
	if dateCheck != nil {
		return 0, errors.New("task date is invalid")
	}

	if task.Date < time.Now().Format("20060102") {
		if task.Repeat == "" {
			task.Date = time.Now().Format("20060102")
		}
		if task.Repeat != "" {
			var err error
			task.Date, err = repeat.NextDate(time.Now(), task.Date, task.Repeat)
			if err != nil {
				return 0, errors.WithMessage(err, "task date is invalid")
			}
		}
	}

	return s.repo.CreateTask(task)
}

func (s *todoService) UpdateTask(task *models.Task) error {
	if task == nil {
		return errors.New("task is nil")
	}

	if task.Id == "" {
		return errors.New("task id is empty")
	}

	_, idCheck := strconv.Atoi(task.Id)
	if idCheck != nil {
		return errors.New("task id is invalid")
	}

	if task.Title == "" {
		return errors.New("task title is empty")
	}

	if task.Date == "" {
		task.Date = time.Now().Format("20060102")
	}
	_, dateCheck := time.Parse("20060102", task.Date)
	if dateCheck != nil {
		return errors.New("task date is invalid")
	}

	if task.Date < time.Now().Format("20060102") {
		if task.Repeat == "" {
			task.Date = time.Now().Format("20060102")
		}
		if task.Repeat != "" {
			var err error
			task.Date, err = repeat.NextDate(time.Now(), task.Date, task.Repeat)
			if err != nil {
				return errors.WithMessage(err, "task date is invalid")
			}
		}
	}

	return s.repo.UpdateTask(task)
}

func (s *todoService) DeleteTask(id int64) error {
	if id == 0 {
		return errors.New("task id is empty")
	}
	return s.repo.DeleteTask(id)
}

func (s *todoService) GetTaskById(id int64) (models.Task, error) {
	return s.repo.GetTaskById(id)
}
