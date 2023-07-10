package main

import (
	"encoding/json"
	"log"
	"net/http"
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

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
