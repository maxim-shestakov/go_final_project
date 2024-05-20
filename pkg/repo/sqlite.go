package repo

import (
	"database/sql"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/maxim-shestakov/final-yandex-project/internal/models"

	"github.com/pkg/errors"
)

type SqliteRepo struct {
	db *sql.DB
}

func NewSqliteRepo(db *sql.DB) *SqliteRepo {
	return &SqliteRepo{
		db: db,
	}
}

func (r *SqliteRepo) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	rows, err := r.db.Query(`SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC LIMIT 20`)
	if err != nil {
		return nil, errors.WithMessage(err, "get tasks: repo select")
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, errors.WithMessage(err, "get tasks: repo scan")
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *SqliteRepo) CreateTask(task *models.Task) (int64, error) {
	res, err := r.db.Exec(`INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, errors.WithMessage(err, "create task: insert")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.WithMessage(err, "create task: get id")
	}

	return id, nil
}

func (r *SqliteRepo) UpdateTask(task *models.Task) error {
	id, err := strconv.ParseInt(task.Id, 10, 64)
	if err != nil {
		return errors.WithMessage(err, "update task: parse id")
	}
	_, err = r.GetTaskById(int64(id))
	if err != nil {
		return errors.New("task not found")
	}
	_, err = r.db.Exec(`
	UPDATE scheduler
	SET date = ?, title = ?, comment = ?, repeat = ?
	WHERE id = ?`, task.Date, task.Title, task.Comment, task.Repeat, int64(id))
	if err != nil {
		return errors.WithMessage(err, "update task: update")
	}

	return nil
}

func (r *SqliteRepo) DeleteTask(id int64) error {
	_, err := r.db.Exec(`DELETE FROM scheduler WHERE id = ?`, id)
	if err != nil {
		return errors.WithMessage(err, "delete task: delete")
	}
	return nil
}

func (r *SqliteRepo) GetTaskById(id int64) (models.Task, error) {
	var task models.Task
	err := r.db.QueryRow(`
	SELECT id, date, title, comment, repeat 
	FROM scheduler WHERE id = ?`, id).
		Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return task, errors.WithMessage(err, "get task by id: select")
	}
	return task, nil
}
