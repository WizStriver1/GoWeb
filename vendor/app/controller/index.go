package controller

import (
	"net/http"
)

// ToIndex displays the home page
func ToIndex() {
	http.FileServer(http.Dir("/static"))
}

// IndexGET displays the home page
func IndexGET(w http.ResponseWriter, r *http.Request) {
	// http.Redirect(w, r, "/static/index.html", 301)
	ToIndex()
}
