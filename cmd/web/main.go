package main

import (
	"log"
	"net/http"

	"github.com/RudyItza/mekah-tell-yuh/internal/app"
	"github.com/RudyItza/mekah-tell-yuh/internal/db"
)

func main() {
	// Initialize the database connection
	database := db.InitDB()
	defer database.Close()

	// Create application instance
	application := app.NewApplication(database)

	// Register routes
	app.RegisterRoutes(application)

	log.Println("Server started on :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
