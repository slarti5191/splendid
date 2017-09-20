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

var templates = template.Must(template.ParseFiles("web/status.html"))

func renderTemplate(w http.ResponseWriter, c *configuration.Config) {
	err := templates.ExecuteTemplate(w, "status.html",
		map[string]interface{}{
			"Config":    c,
			"TimeSince": time.Since(start),
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RunTheServer() {
	start = time.Now()

	http.HandleFunc("/", testHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", nil))
}
