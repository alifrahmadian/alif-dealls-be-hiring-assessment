package db

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

type DBConfig struct {
	Host string
	Port string
	Username string
	Password string
	DBName string
}

func Connect(dbConfig DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.DBName)

	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return DB, nil
}