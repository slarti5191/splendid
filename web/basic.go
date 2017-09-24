package web

import (
	"github.com/slarti5191/splendid/configuration"
	"html/template"
	"log"
	"net/http"
	"time"
)

var start time.Time

func testHandler(w http.ResponseWriter, r *http.Request) {
	c := configuration.GetConfig()
	renderTemplate(w, c)
}

var templates = template.Must(template.ParseFiles(
	"resources/templates/layouts/master.html",
	"resources/templates/pages/status.html",
))

func renderTemplate(w http.ResponseWriter, c *configuration.Config) {
	err := templates.ExecuteTemplate(w, "master.html",
		map[string]interface{}{
			"Config":    c,
			"TimeSince": time.Since(start),
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RunTheServer kicks off the http listener.
func RunTheServer() {
	start = time.Now()

	http.HandleFunc("/", testHandler)
	http.Handle("/assets/", http.FileServer(http.Dir("resources/static")))

	// Bind to a port and pass our router in
	log.Println("> Starting web server on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
