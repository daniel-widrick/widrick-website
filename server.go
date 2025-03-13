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
		filePath := filepath.Join("./", r.URL.Path)
		normalizedPath := strings.ReplaceAll(strings.TrimPrefix(filePath, "/"), "\\", "/")
		if strings.HasPrefix(normalizedPath, "portfolio/") {
			validLocations := []string{"software", "systems", "security"}
			isValid := false
			for _, loc := range validLocations {
				if strings.HasPrefix(normalizedPath, "portfolio/"+loc) {
					isValid = true
					break
				}
			}
			if !isValid {
				http.NotFound(w, r)
				return
			}
			filePath = "./index.html"
		}
		log.Printf("%s  :: Serving file: %s (%s)\n", r.RemoteAddr, filePath, r.URL.Path)
		http.ServeFile(w, r, filePath)
	})

	log.Println("Starting server on :6780")
	log.Fatalln(http.ListenAndServe(":6780", mux))
}
