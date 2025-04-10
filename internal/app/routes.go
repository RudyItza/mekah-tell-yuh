package app

import "net/http"

// RegisterRoutes sets up the routes for the application
func RegisterRoutes(app *Application) {
	// Route to view the home page as /home
	http.HandleFunc("/home", app.homeHandler)

	// Route to submit a new story
	http.HandleFunc("/submit", app.submitStoryHandler)

	// Route to view all stories
	http.HandleFunc("/stories", app.viewStoriesHandler)

	// Route to edit a story (note the id extraction from the URL path)
	http.HandleFunc("/edit/", app.editStoryHandler) // /edit/{id}

	// Route to delete a story (note the id extraction from the URL path)
	http.HandleFunc("/delete/", app.deleteStoryHandler) // /delete/{id}
}
