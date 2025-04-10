package app

import (
	"database/sql"

	"github.com/RudyItza/mekah-tell-yuh/internal/data" // Import the data package for StoryModel
)

// Application holds the necessary components for the application, including the DB and StoryModel.
type Application struct {
	DB         *sql.DB
	StoryModel *data.StoryModel // Add StoryModel to the Application struct
}

// NewApplication creates a new application with the provided DB connection and StoryModel
func NewApplication(db *sql.DB, storyModel *data.StoryModel) *Application {
	return &Application{
		DB:         db,
		StoryModel: storyModel, // Initialize StoryModel
	}
}
