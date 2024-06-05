package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

  "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type AppDatabase struct {
	*sqlx.DB

	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
	SslMode  string `json:"sslmode"`
}

func NewAppDatabase() *AppDatabase {
	return &AppDatabase{}
}

func (db *AppDatabase) LoadConfig(path string) error {
	buf, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, &db)
}

func (db *AppDatabase) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s",
		db.Host, db.Port, db.User, db.Password)

	if db.DbName != "" {
		psqlInfo += fmt.Sprintf(" dbname=%s", db.DbName)
	}

	if db.SslMode != "" {
		psqlInfo += fmt.Sprintf(" sslmode=%s", db.SslMode)
	}

	dbClient, err := sqlx.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.DB = dbClient
	err = db.DB.Ping()
	return err
}
