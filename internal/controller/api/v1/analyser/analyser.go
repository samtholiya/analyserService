package analyser

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Analyser struct {
	Crawler Crawler
}

func (a Analyser) Post(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
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

	url, ok := reqData["data"]
	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	responseData, err := a.Crawler.FromURL(url)
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
