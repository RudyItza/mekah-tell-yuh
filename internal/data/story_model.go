package data

import (
	"database/sql"

	"github.com/RudyItza/mekah-tell-yuh/internal/app"
)

// StoryModel defines methods to interact with the 'stories' table
type StoryModel struct {
	DB *sql.DB
}

// NewStoryModel creates a new StoryModel instance
func NewStoryModel(db *sql.DB) *StoryModel {
	return &StoryModel{DB: db}
}

// Insert inserts a new story into the database
func (m *StoryModel) Insert(title, content string) (int, error) {
	var id int
	err := m.DB.QueryRow("INSERT INTO stories (title, content) VALUES ($1, $2) RETURNING id", title, content).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetAll retrieves all stories from the database
func (m *StoryModel) GetAll() ([]app.Story, error) {
	rows, err := m.DB.Query("SELECT id, title, content FROM stories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stories []app.Story
	for rows.Next() {
		var s app.Story
		if err := rows.Scan(&s.ID, &s.Title, &s.Content); err != nil {
			return nil, err
		}
		stories = append(stories, s)
	}

	return stories, nil
}

// GetByID retrieves a story by its ID
func (m *StoryModel) GetByID(id int) (*app.Story, error) {
	var story app.Story
	err := m.DB.QueryRow("SELECT id, title, content FROM stories WHERE id = $1", id).Scan(&story.ID, &story.Title, &story.Content)
	if err != nil {
		return nil, err
	}
	return &story, nil
}

// Update updates a story's title and content
func (m *StoryModel) Update(id int, title, content string) error {
	_, err := m.DB.Exec("UPDATE stories SET title = $1, content = $2 WHERE id = $3", title, content, id)
	return err
}

// Delete deletes a story by its ID
func (m *StoryModel) Delete(id int) error {
	_, err := m.DB.Exec("DELETE FROM stories WHERE id = $1", id)
	return err
}
