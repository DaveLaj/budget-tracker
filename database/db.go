package database

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// global varz
var (
	BudgetTrackerPool *sql.DB
)

type Config struct {
	DB_USER     string
	DB_HOST     string
	DB_PASSWORD string
	DB_PORT     int
}

func InitializeDB(cfg Config, dbName string) error {
	connection, err := CreateDBConnection(cfg, dbName)
	if err != nil {
		return fmt.Errorf("error creating database connection: %w", err)
	}

	BudgetTrackerPool = connection
	return nil
}

func CreateDBConnection(cfg Config, dbName string) (*sql.DB, error) {
	var db *sql.DB
	var err error

	dbAddr := fmt.Sprintf("%s:%d", cfg.DB_HOST, cfg.DB_PORT)
	dbCfg := mysql.Config{
		User:                 cfg.DB_USER,
		Passwd:               cfg.DB_PASSWORD,
		Net:                  "tcp",
		Addr:                 dbAddr,
		DBName:               dbName,
		AllowNativePasswords: true,
	}

	db, err = sql.Open("mysql", dbCfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to database")
	return db, nil
}
