package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	appName         = "svc-api"
	appVersion      = "0.0.1"
	httpDefaultPort = "8080"
	topicPubNotify  = "notify"
	projectID       = "pontos-filhos"
)

func main() {
	// service initialization
	log.SetFlags(log.Lshortfile)

	log.Println("Starting", appName, "version", appVersion)

	port := os.Getenv("PORT")
	if port == "" {
		port = httpDefaultPort
		log.Printf("Service: %s. Defaulting to port %s", appName, port)
	}

	http.HandleFunc("/", indexHandler)

	// Start web server
	log.Printf("Service: %s. Listening on port %s", appName, port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}