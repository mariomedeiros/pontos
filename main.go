package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

	

	// HTTP Server initialization
	// define all the routes for the HTTP server.
	//   The implementation is done on the "handler*.go" files
	router := mux.NewRouter()
	// version
	router.HandleFunc("/api/version", getVersion).Methods("GET")
	// Callpoints
	//router.HandleFunc("/api/pontos", getPontosSoma).Methods("GET")
	//router.HandleFunc("/api/pontos", postPontos).Methods("POST")
	
	// Start web server
	log.Printf("Service: %s. Listening on port %s", appName, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

//processError generic error processing method to fill in default HTTP content
func processError(e error, w http.ResponseWriter, httpCode int, status string, detail string) {
	log.Println(e)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, `{"status":"%s", "description":"%s", "fullError":"%s"}`, status, detail, e.Error())
}

// getVersion get version
func getVersion(w http.ResponseWriter, r *http.Request) {
	log.Println("[/version:GET] Requested api version. " + appVersion)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, fmt.Sprintf(`{"service": "%s", "version": "%s"}`, appName, appVersion))
}

