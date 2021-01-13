package analyser

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Analyser struct {
	Crawler Crawler
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (a Analyser) Post(rw http.ResponseWriter, req *http.Request) {
	setupResponse(&rw, req)
	if (*req).Method == "OPTIONS" {
		return
	}
	if req.Method != "POST" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	reqData := make(map[string]string)
	err := json.NewDecoder(req.Body).Decode(&reqData)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	urlString, ok := reqData["data"]
	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, err := url.ParseRequestURI(urlString); err != nil {
		rw.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	responseData, err := a.Crawler.FromURL(urlString)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "no such host") {
			rw.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(responseData)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
	_, err = rw.Write(data)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
