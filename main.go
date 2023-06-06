package main

import (
	"api/handler"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler.Read)
	http.ListenAndServe(":8000", nil)
}
