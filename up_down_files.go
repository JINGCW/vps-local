package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

const (
	max_upload_size = 5 * 1024 * 1024
	upload_path     = "uploads/"
)

func main() {
	http.HandleFunc("/upload", upload_file_handler())

	fs := http.FileServer(http.Dir(upload_path))
	http.Handle("/files/", http.StripPrefix("/files/", fs))

	log.Print("\n###server started on http://0.0.0.0:8080\n###" +
		"use /upload for uploading files and /files/{file_name} for downloading files.")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func upload_file_handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//validate file size
		r.Body = http.MaxBytesReader(w, r.Body, max_upload_size)
		if err := r.ParseMultipartForm(max_upload_size); err != nil {
			render_error(w, "file too big", http.StatusBadRequest)
			return
		}
		//parse and validate file and post parameters
		file, _, err := r.FormFile("uploadFile")
		if err != nil {
			render_error(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		defer file.Close()
		file_bytes, err := ioutil.ReadAll(file)
		if err != nil {
			render_error(w, "INVALID_FILE", 400)
			return
		}
		//check file type, detect content type only needs the first 512 bytes
		detected_file_type := http.DetectContentType(file_bytes)
		switch detected_file_type {
		case "image/jpeg", "image/jpg", "image/gif", "image/png", "application/pdf":
			break
		default:
			render_error(w, "INVALID_FILE_TYPE", 400)
			return
		}

		file_name := rand_token(12)
		file_endings, err := mime.ExtensionsByType(detected_file_type)
		//fmt.Println(file_endings)
		if err != nil {
			render_error(w, "CANT_READ_FILE_TYPE", 400)
			return
		}
		//fmt.Println("new file path name is :",file_name+file_endings[0])
		new_path := filepath.Join(upload_path, file_name+file_endings[0])
		fmt.Printf("FileType: %s, File: %s\n", detected_file_type, new_path)
		//write file
		_, err = os.Stat(new_path)
		if os.IsNotExist(err) {
			os.Mkdir(filepath.Dir(new_path), os.ModePerm)
		}
		new_file, err := os.Create(new_path)
		if err != nil {
			render_error(w, "CANT_WRITE_FILE", 400)
			return
		}
		defer new_file.Close()
		if _, err := new_file.Write(file_bytes); err != nil || new_file.Close() != nil {
			render_error(w, "CANT_WRITE_FILE", 400)
			return
		}
		w.Write([]byte("SUCCESS"))
	}
}

func render_error(w http.ResponseWriter, msg string, status_code int) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte(msg))
	if err != nil {
		fmt.Println("render_error error.")
		os.Exit(1)
	}
}

func rand_token(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("rand_token error.")
		os.Exit(1)
	}
	return fmt.Sprintf("%x", b)
}
