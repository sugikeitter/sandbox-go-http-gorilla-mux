package main

import (
	"encoding/json"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		nt := time.Now()
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello", "time": nt.Format("2006/01/02 15:04:05.000")})
	})
	r.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		reqPathVars := mux.Vars(r)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello, " + reqPathVars["name"]})
	})

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
