package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type DBInstance struct {
	initializer func() interface{}
	instance    interface{}
	once        sync.Once
}

func (i *DBInstance) Instance() interface{} {
	i.once.Do(func() {
		i.instance = i.initializer()
	})

	return i.instance
}

func dbInit() interface{} {
	maxRetries := 5
	var db *sql.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName))
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		if err := db.Ping(); err == nil {
			return db
		}

		db.Close()

		if i < maxRetries-1 {
			time.Sleep(5 * time.Second)
		}
	}

	return nil
}

func Connect() *sql.DB {
	dbInstance := &DBInstance{initializer: dbInit}
	return dbInstance.Instance().(*sql.DB)
}
