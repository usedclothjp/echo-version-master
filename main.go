package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	bind := os.Getenv("BIND")
	if bind == "" {
		bind = ":8080"
	}

	appVersion := os.Getenv("APP_VERSION")
	if appVersion == "" {
		appVersion = "0.0.1"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "APP_VERSION=%s", appVersion)
		log.Printf("Receive GET request path = %s\n", r.RequestURI)
	})
	http.ListenAndServe(bind, nil)
}
