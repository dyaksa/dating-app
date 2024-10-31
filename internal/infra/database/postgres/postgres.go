package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

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
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName))
	if err != nil {
		os.Exit(1)
		db.Close()
	}

	if err := db.Ping(); err != nil {
		os.Exit(1)
		db.Close()
	}

	return db
}

func Connect() *sql.DB {
	dbInstance := &DBInstance{initializer: dbInit}
	return dbInstance.Instance().(*sql.DB)
}
