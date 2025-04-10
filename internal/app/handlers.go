package app

import (
	"net/http"
	"strconv"
)

func (app *Application) homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.tmpl", nil)
}

func (app *Application) submitStoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")

		_, err := app.DB.Exec("INSERT INTO stories (title, content) VALUES ($1, $2)", title, content)
		if err != nil {
			http.Error(w, "Failed to submit story.", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/stories", http.StatusSeeOther)
		return
	}
	renderTemplate(w, "submit_story.tmpl", nil)
}

func (app *Application) viewStoriesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := app.DB.Query("SELECT id, title, content FROM stories")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var stories []Story
	for rows.Next() {
		var s Story
		if err := rows.Scan(&s.ID, &s.Title, &s.Content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		stories = append(stories, s)
	}

	renderTemplate(w, "view_story.tmpl", stories)
}

func (app *Application) editStoryHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid story ID.", http.StatusBadRequest)
		return
	}

	if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")

		_, err := app.DB.Exec("UPDATE stories SET title = $1, content = $2 WHERE id = $3", title, content, id)
		if err != nil {
			http.Error(w, "Failed to update story.", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/stories", http.StatusSeeOther)
		return
	}

	var story Story
	err = app.DB.QueryRow("SELECT id, title, content FROM stories WHERE id = $1", id).Scan(&story.ID, &story.Title, &story.Content)
	if err != nil {
		http.Error(w, "Story not found.", http.StatusNotFound)
		return
	}

	renderTemplate(w, "edit_story.tmpl", story)
}

func (app *Application) deleteStoryHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid story ID.", http.StatusBadRequest)
		return
	}

	_, err = app.DB.Exec("DELETE FROM stories WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Failed to delete story.", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/stories", http.StatusSeeOther)
}
