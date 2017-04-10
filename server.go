package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func allowedFile(s string) bool {
	allowed := []string{".txt", ".pdf", ".png", ".jpg", ".jpeg", ".gif", ".mov", ".mp4", ".mts"}
	for _, ext := range allowed {
		if strings.HasSuffix(strings.ToLower(s), ext) {
			return true
		}
	}
	return false
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2 * 1024 * 1024)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid request", 400)
		return
	}
	defer file.Close()
	if len(handler.Filename) > 0 && allowedFile(handler.Filename) {
		f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "invalid request", 400)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		w.WriteHeader(http.StatusCreated)
		return
	}
	http.Error(w, "type not allowed", 415)
	return
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	uploadDir := path.Join(pwd, "uploads")
	fmt.Println("Upload Dir:", uploadDir)

	http.Handle("/", http.FileServer(http.Dir(pwd)))
	http.Handle("/uploads/", http.StripPrefix("/uploads", http.FileServer(http.Dir(uploadDir))))
	http.HandleFunc("/upload", handleUpload)
	err = http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
