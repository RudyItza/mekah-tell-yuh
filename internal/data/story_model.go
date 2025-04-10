package data

import (
	"database/sql"
	"time"
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
func (m *StoryModel) Insert(title, content, language, location, category string, userID int64) (int64, error) {
	var id int64
	query := `INSERT INTO stories (title, content, language, location, category, user_id, created_at, updated_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := m.DB.QueryRow(query, title, content, language, location, category, userID, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetAll retrieves all stories from the database
func (m *StoryModel) GetAll() ([]Story, error) { // Change `app.Story` to `Story` directly
	rows, err := m.DB.Query("SELECT id, title, content, language, location, category, user_id, created_at, updated_at FROM stories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stories []Story
	for rows.Next() {
		var s Story
		if err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Language, &s.Location, &s.Category, &s.UserID, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		stories = append(stories, s)
	}

	return stories, nil
}

// Update updates an existing story in the database
func (m *StoryModel) Update(id int, title, content string) error {
	query := `UPDATE stories SET title = $1, content = $2, updated_at = $3 WHERE id = $4`
	_, err := m.DB.Exec(query, title, content, time.Now(), id)
	return err
}

// GetByID retrieves a story by its ID
func (m *StoryModel) GetByID(id int) (*Story, error) { // Change `app.Story` to `Story` directly
	query := `SELECT id, title, content, language, location, category, user_id, created_at, updated_at FROM stories WHERE id = $1`
	row := m.DB.QueryRow(query, id)

	var s Story
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Language, &s.Location, &s.Category, &s.UserID, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// Delete deletes a story by its ID
func (m *StoryModel) Delete(id int) error {
	query := `DELETE FROM stories WHERE id = $1`
	_, err := m.DB.Exec(query, id)
	return err
}
