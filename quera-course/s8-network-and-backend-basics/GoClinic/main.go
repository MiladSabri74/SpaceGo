package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := Response{
		Status: "up",
	}
	json.NewEncoder(w).Encode(response)

}

func manageAPIs() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/health", healthHandler)
}

func StartServer(port string) {
	// TODO: start the HTTP server with /ping and /health endpoints
	manageAPIs()

	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
