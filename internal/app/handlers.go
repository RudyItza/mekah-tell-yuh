package app

import (
	"log"
	"net/http"
	"strconv"
)

// SubmitStoryHandler handles the submission of a new story
func (app *Application) submitStoryHandler(w http.ResponseWriter, r *http.Request) {
	// Log the user ID from the session for debugging
	log.Println("User ID from session:", app.Session.UserID)

	if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")
		// location := r.FormValue("location") // Removed as it is not used
		// Removed unused variable 'category'
		userID, err := strconv.Atoi(r.FormValue("user_id"))
		if err != nil {
			http.Error(w, "Invalid user ID.", http.StatusBadRequest)
			return
		}

		// Validate title and content
		if err := ValidateStoryInput(title, content); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Insert the story into the database
		language := ""
		location := ""
		category := ""
		_, err = app.StoryModel.Insert(title, content, language, location, category, int64(userID))
		if err != nil {
			http.Error(w, "Failed to submit story.", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/stories", http.StatusSeeOther)
		return
	}

	renderTemplate(w, "submit_story.tmpl", nil)
}

// ViewStoriesHandler retrieves and displays all stories
func (app *Application) viewStoriesHandler(w http.ResponseWriter, r *http.Request) {
	stories, err := app.StoryModel.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderTemplate(w, "view_story.tmpl", stories)
}

// EditStoryHandler allows editing of an existing story
func (app *Application) editStoryHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the story ID from the URL query string
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid story ID.", http.StatusBadRequest)
		return
	}

	if r.Method == "POST" {
		// Retrieve the updated form values
		title := r.FormValue("title")
		content := r.FormValue("content")
		// Removed unused variable 'language'
		// Removed unused variable 'category'

		// Validate title and content
		if err := ValidateStoryInput(title, content); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update the story in the database with the additional fields
		err := app.StoryModel.Update(id, title, content)
		if err != nil {
			http.Error(w, "Failed to update story.", http.StatusInternalServerError)
			return
		}

		// Redirect to the list of stories after successful update
		http.Redirect(w, r, "/stories", http.StatusSeeOther)
		return
	}

	// Retrieve the story to be edited
	story, err := app.StoryModel.GetByID(id)
	if err != nil {
		http.Error(w, "Story not found.", http.StatusNotFound)
		return
	}

	// Render the edit story template with the story data
	renderTemplate(w, "edit_story.tmpl", story)
}

// DeleteStoryHandler deletes a story by ID
func (app *Application) deleteStoryHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid story ID.", http.StatusBadRequest)
		return
	}

	err = app.StoryModel.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete story.", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/stories", http.StatusSeeOther)
}

// homeHandler serves the home page
func (app *Application) homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home page accessed")         // Debugging log
	http.ServeFile(w, r, "ui/html/home.tmpl") // Serve the home.html file
}
