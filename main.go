package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
)

type Answer struct {
    Name    string `json:"nome"`
    Time    string `json:"horario"`
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	formattedTime := time.Now().UTC().Format(time.RFC3339)

	response := Answer {
		Name: "Projeto Korp",
		Time: formattedTime,
	}

	json.NewEncoder(w).Encode(response)

}

func main () {
	http.HandleFunc("/projeto-korp", handlerFunc)
	
	fmt.Println("Server is running at http://localhost:8080/projeto-korp")
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
