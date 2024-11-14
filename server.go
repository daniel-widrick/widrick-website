package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main(){
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
		filePath := filepath.Join("./",r.URL.Path)
		log.Printf("%s  :: Serving file: %s\n",r.RemoteAddr,filePath)
		http.ServeFile(w, r, filePath)
	})

	log.Println("Starting server on :6780")
	log.Fatalln(http.ListenAndServe(":6780",mux))
}	
