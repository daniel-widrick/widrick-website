package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func main(){
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
		filePath := filepath.Join("./",r.URL.Path)
		switch strings.TrimPrefix(filePath,"/") {
		case "portfolio/software":
			filePath = "./index.html"
		case "portfolio/systems":
			filePath = "./index.html"
		case "portfolio/security":
			filePath = "./index.html"
		}
		log.Printf("%s  :: Serving file: %s\n",r.RemoteAddr,filePath)
		http.ServeFile(w, r, filePath)
	})

	log.Println("Starting server on :6780")
	log.Fatalln(http.ListenAndServe(":6780",mux))
}	
