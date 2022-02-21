package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const maxUploadSize = 20 * 1024 * 1024 // 20 mb
const mdUploadPath = "markdown"
const imgUploadPath = "images"
const imgUrl = "http://img.chion.tech/"

func main() {
	// routing
	http.HandleFunc("/upload", uploadFileHandler())
	fs := http.FileServer(http.Dir(imgUploadPath))
	http.Handle("/", http.StripPrefix("/images", fs))
	log.Print("Server started on localhost:8081")

	//Port services
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func uploadFileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			t, _ := template.ParseFiles("upload.gtpl")
			t.Execute(w, nil)
			return
		}
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			fmt.Printf("Could not parse multipart form: %v\n", err)
			renderError(w, "CANT_PARSE_FORM")
			return
		}

		// parse and validate file and post parameters
		file, fileHeader, err := r.FormFile("uploadFile")
		if err != nil {
			renderError(w, "INVALID_FILE")
			return
		}
		defer file.Close()
		// Get and print out file size
		fileSize := fileHeader.Size
		//fmt.Printf("File size (bytes): %v\n", fileSize)
		// validate file size
		if fileSize > maxUploadSize {
			renderError(w, "FILE_TOO_BIG")
			return
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID_FILE")
			return
		}

		// check file type, Affectionate only needs the first 512 bytes
		detectedFileType := http.DetectContentType(fileBytes)
		switch detectedFileType {
		case "image/jpeg", "image/jpg":
		case "image/gif", "image/png":
		case "text/plain; charset=utf-8":
		case "application/octet-stream", "application/pdf":
			break
		default:
			renderError(w, "INVALID_FILE_TYPE")
			return
		}
		fileEndings, err := mime.ExtensionsByType(detectedFileType)
		if err != nil {
			renderError(w, "CANT_READ_FILE_TYPE")
			return
		}
		var newPath string
		var fileName string
		if strings.EqualFold(fileEndings[0], ".conf") {
			// 由于最后需要编译成为二进制文件，所以不获取当前路径
			fileName = fileHeader.Filename
			newPath = filepath.Join(mdUploadPath, fileName)
			fmt.Printf("File: %s\n", newPath)
		} else {
			fileName = time.Now().Format("20060102150405") + fileEndings[0]
			newPath = filepath.Join(imgUploadPath, fileName)
			fmt.Printf("imgUrl: %s, File: %s\n", imgUrl+fileName, newPath)
		}
		//fmt.Printf("FileType: %s, File: %s\n", detectedFileType, newPath)

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			renderError(w, "CANT_WRITE_FILE")
			return
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE")
			return
		}
		if strings.EqualFold(fileEndings[0], ".conf") {
			w.Write([]byte("<b>File upload Successful! File：</b>" + imgUrl + fileName))
		} else {
			w.Write([]byte(imgUrl + fileName))
		}
	}
}

func renderError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}
