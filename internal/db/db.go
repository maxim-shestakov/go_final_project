package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// type DBConfig struct {
// 	User     string
// 	Password string
// 	Host     string
// 	Port     string
// 	Database string
// }

// func New(config *DBConfig) (*sql.DB, error) {
// 	psqlInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Database)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

func InitDB() (*sql.DB, error) {
	// appPath, err := os.Executable()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var dbFile string
	// dbpath := os.Getenv("TODO_DBFILE")
	// if len(dbpath) == 0 {
	// 	// dbFile = filepath.Join(filepath.Dir(appPath), "scheduler.db")
	// 	dbFile = "scheduler.db"
	// } else {
	// 	dbFile = filepath.Join(dbpath, "scheduler.db")
	// }

	// _, err := os.Stat(dbFile)

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file := os.Getenv("TODO_DBFILE")

	if len(file) == 0 {
		file = "scheduler.db"
	}

	dbFile := filepath.Join(currentDir, file)
	_, err = os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	fmt.Println("dbFile: ", dbFile)
	// если install равен true, после открытия БД требуется выполнить
	// sql-запрос с CREATE TABLE и CREATE INDEX

	if install {
		db, err := sql.Open("sqlite3", dbFile)
		if err != nil {
			return nil, err
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS scheduler (
			id INTEGER PRIMARY KEY,
			date char(8),
			title varchar(255),
			comment TEXT,
			repeat varchar(128)
		)`)
		if err != nil {
			return nil, err
		}
		_, err = db.Exec(`CREATE INDEX  IF NOT EXISTS idx_scheduler_date ON scheduler (date)`)
		if err != nil {
			return nil, err
		}
		return db, nil
	}

	return sql.Open("sqlite3", dbFile)
}
