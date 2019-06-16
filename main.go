package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	router := NewRouter()
	router.PathPrefix("/").HandlerFunc(IndexHandler)
	http.ListenAndServe(":8888", router)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path == "" {
		path = "index.html"
	}
	log.Println(path)

	data, err := ioutil.ReadFile("docs/" + string(path))

	if err == nil {
		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else if strings.HasSuffix(path, ".jpg") {
			contentType = "image/jpeg"
		} else {
			contentType = "text/plain"
		}

		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 this didnt work" + http.StatusText(404)))
	}
}