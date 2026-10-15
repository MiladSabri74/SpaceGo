package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Ticket struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Ride string `json:"ride"`
}

var Tickets map[int]Ticket
var AutoIncrimentID int

func ticketHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		ride := r.URL.Query().Get("ride")
		var result []Ticket
		for _, t := range Tickets {
			if t.Ride == ride || ride == "" {
				result = append(result, t)
			}
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		j, _ := json.MarshalIndent(&result, "", "  ")
		w.Write(j)

	} else if r.Method == http.MethodPost {
		var t Ticket
		json.NewDecoder(r.Body).Decode(&t)
		t.ID = AutoIncrimentID
		AutoIncrimentID++
		Tickets[t.ID] = t
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&t)

	} else if r.Method == http.MethodDelete {
		w.Header().Set("content-type", "application/json")

		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		if _, ok := Tickets[id]; ok {
			fmt.Println("exist")
			delete(Tickets, id)
			w.WriteHeader(http.StatusOK)
		} else {
			fmt.Println("not exist")
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func StartServer(port string) {
	// TODO: start the HTTP server with /tickets CRUD endpoints
	Tickets = make(map[int]Ticket)
	AutoIncrimentID = 1
	http.HandleFunc("/tickets", ticketHandler)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
