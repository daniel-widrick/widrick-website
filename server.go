package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// RFC 2324 §2.3.2: a teapot MUST refuse to brew coffee, and SHOULD return 418.
// This one refuses politely, then offers something it can actually make.
const teapotBody = `418 I'm a teapot

I can't brew coffee - I'm a teapot. But the pumps are primed, so here's
what I *can* make you:

    The House Pour
    ------------------------------
    2 parts  vodka
    3 parts  Ocean Spray
    2 parts  Sprite Zero
    Serve in a mason jar.

- Daniel
`

// withHireMe stamps a note on every response, for anyone reading the headers.
func withHireMe(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Hire-Me", "You read response headers for fun - we'd get along. https://widrick.net")
		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()

	// GET /coffee - the teapot cannot brew coffee (HTTP 418). It has a recipe instead.
	mux.HandleFunc("/coffee", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s  :: 418 teapot (%s)\n", r.RemoteAddr, r.URL.Path)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte(teapotBody))
	})

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
	log.Fatalln(http.ListenAndServe(":6780", withHireMe(mux)))
}
