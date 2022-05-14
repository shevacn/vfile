package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var uploadUiHtml string

func uploadUi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, uploadUiHtml)
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	f, err := os.OpenFile("/tmp/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintln(w, "upload ok!")
}

func main() {
	b, _ := os.ReadFile("uploadUi.html")
	uploadUiHtml = string(b)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/uploadUi", uploadUi)
	http.Handle("/ls/", http.StripPrefix("/ls/", http.FileServer(http.Dir("/"))))

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
