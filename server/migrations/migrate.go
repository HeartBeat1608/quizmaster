package main

import (
	"log"
	"os"
	"quizmaster-server/pkg/database"
	"strings"
)

type ParsedStatement struct {
  Stmt string
}

func ParseStatements(raw string) []ParsedStatement {
  var statements []ParsedStatement

  parts := strings.Split(raw, "--")
  for _, part := range parts {
    part = strings.Trim(part, "\n")
    if part != "" {
      statements = append(statements, ParsedStatement{part})
    }
  }

  return statements
}

func Migrate(db *database.AppDatabase) error {
	entires, err := os.ReadDir(".")
	if err != nil {
		return err
	}

	var files []string

	for _, entry := range entires {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".sql") {
			files = append(files, entry.Name())
		}
	}

	for _, file := range files {
		log.Println("Migrating file: ", file)
		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		rawQuery := string(data)
    queries := ParseStatements(rawQuery)

    for _, query := range queries {
      _, err := db.Exec(query.Stmt)
      if err != nil {
        return err
      }
    }
	}

	return nil
}

func main() {
	db := database.NewAppDatabase()
	err := db.LoadConfig("../config/database.json")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Printf("Connected to database: %s:%d", db.Host, db.Port)

	if err = Migrate(db); err != nil {
		log.Fatal(err)
	}

	log.Println("Database Migrations done")
}
