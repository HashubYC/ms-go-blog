package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
)

type IndexData struct {
	Title string
	Desc  string
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	var indexData IndexData
	indexData.Title = "This Title"
	indexData.Desc = "This Desc"
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
	// w.Write([]byte("hello world"))
}
func indexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "This Title"
	indexData.Desc = "This Desc"
	t := template.New("index.html")
	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w, indexData)
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
