package app

import "net/http"

func RegisterRoutes(app *Application) {
	http.HandleFunc("/", app.homeHandler)
	http.HandleFunc("/submit", app.submitStoryHandler)
	http.HandleFunc("/stories", app.viewStoriesHandler)
	http.HandleFunc("/edit", app.editStoryHandler)     // New route for editing stories
	http.HandleFunc("/delete", app.deleteStoryHandler) // New route for deleting stories
}
