package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	service "go-http-rest/internal/service"
)

const (
	indexFile     = "index.html"
	formFileField = "myFile"
	replyHeader   = "text/html; charset=utf-8"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, indexFile)

}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "only POST requests allowed", http.StatusBadRequest)
		log.Printf("got not a POST request")
		return
	}
	fileUploaded, _, err := r.FormFile(formFileField)
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		log.Printf("ParseForm error: %v", err)
		return
	}
	filename := time.Now().UTC().String()
	newFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, fileUploaded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	output, err := service.ConvertData(string(data))
	if err != nil {
		log.Fatalf("unable to convert data: %v", err)
	}
	w.Header().Set("Content-Type", replyHeader)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}
