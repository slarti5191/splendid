package web

import (
	"github.com/slarti5191/splendid/configuration"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

var start time.Time

var rootTemplate = template.Must(template.ParseFiles(
	"resources/templates/layouts/master.html",
	"resources/templates/pages/status.html",
))
var deviceTemplate = template.Must(template.ParseFiles(
	"resources/templates/layouts/master.html",
	"resources/templates/pages/device.html",
))

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Execute and display the page.
	err := rootTemplate.ExecuteTemplate(w, "master.html",
		map[string]interface{}{
			"Config":    configuration.GetConfig(),
			"TimeSince": time.Since(start),
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func deviceHandler(w http.ResponseWriter, r *http.Request) {
	// Make sure we're on /device/xyz and not a subfolder.
	if path.Dir(r.URL.Path) != "/device" {
		http.NotFound(w, r)
		return
	}

	// Name of target device based on url path.
	name := path.Base(r.URL.Path)
	c := configuration.GetConfig()

	var targetDevice *configuration.DeviceConfig
	for _, device := range c.Devices {
		if device.Name == name {
			targetDevice = &device
			break
		}
	}
	if targetDevice == nil {
		http.NotFound(w, r)
		return
	}

	// Execute and display the page.
	err := deviceTemplate.ExecuteTemplate(w, "master.html",
		map[string]interface{}{
			"Config":    configuration.GetConfig(),
			"TimeSince": time.Since(start),
			"Device":    targetDevice,
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RunTheServer kicks off the http listener.
func RunTheServer() {
	start = time.Now()

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/device/", deviceHandler)
	http.Handle("/assets/", http.FileServer(http.Dir("resources/static")))

	// Bind to a port and pass our router in
	log.Println("> Starting web server on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
