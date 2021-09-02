package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	usageMsg := `Usage:
	<command> <address> <port>
`
	if len(os.Args) != 3 {
		fmt.Printf(usageMsg)
		os.Exit(1)
	}

	addr := os.Args[1]
	port := os.Args[2]

	if addr == "" || port == "" {
		fmt.Printf(usageMsg)
		os.Exit(1)
	}

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
		Addr:    addr + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
