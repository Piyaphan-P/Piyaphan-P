package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Hello struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hello := Hello{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hello)
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
