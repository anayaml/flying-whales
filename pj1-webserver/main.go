package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Print(w, "This is the home page.")
}

func httpMethodHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Get request")
	case "POST":
		fmt.Fprintf(w, "Post request")
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

type Response struct {
	Message string `json:"message"`
}

func jsonResponseHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{"Hello, JSON!"}
	json.NewEncoder(w).Encode(response)
}

func handleRequests() {
	http.HandleFunc("/", getHomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
