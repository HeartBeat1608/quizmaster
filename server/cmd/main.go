package main

import (
	"log"
	"quizmaster-server/core"
	"quizmaster-server/pkg/database"
)

func main() {
	app := core.NewAppServer(":5000")

	db := database.NewAppDatabase()
	if err := db.LoadConfig("config/database.json"); err != nil {
		log.Fatalf("Failed to load database config: %v", err)
	}
	if err := db.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	app.SetupMiddlewares()
	app.SetupRoutes(db)

	log.Printf("Server listening on %s", app.ListenAddr())
	log.Fatal(app.Serve())
}
