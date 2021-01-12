package main

import (
	"log"
	"net/http"
	"os"

	"github.com/samtholiya/analyserService/internal/controller/api/v1/analyser"
	"github.com/samtholiya/analyserService/internal/service/analyser/html"
)

func main() {
	// Generally I would divide the controllers
	// Use interface etc
	// but as there is just one api keep it simple silly
	mux := http.NewServeMux()
	crawl := &html.Crawler{}
	analyser := analyser.Analyser{
		Crawler: crawl,
	}
	mux.HandleFunc("/api/v1/analyser/html", analyser.Post)
	log.Println("Server is starting at port 80")
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
