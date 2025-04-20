package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join("/app", r.URL.Path)
		log.Printf("%s  :: Request: %s (%s)\n", r.RemoteAddr, filePath, r.URL.Path)
		if strings.HasPrefix(filePath, "/app/portfolio") {
			validLocations := []string{"software", "systems", "security"}
			for _, loc := range validLocations {
				locPath := filepath.Join("/app/portfolio", loc)
				if strings.HasPrefix(filePath, locPath) {
					filePath = "/app/index.html"
					break
				}
			}
		}
		if filePath == "/app" {
			filePath = "/app/index.html"
		}
		log.Printf("%s  :: Serving file: %s (%s)\n", r.RemoteAddr, filePath, r.URL.Path)
		http.ServeFile(w, r, filePath)
	})

	log.Println("Starting server on :6780")
	log.Fatalln(http.ListenAndServe(":6780", mux))
}
