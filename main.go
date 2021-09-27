package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var counter = 0

func main() {
	fmt.Printf("start!")

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
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		nt := time.Now()
		counter++
		res := fmt.Sprintf(
			`<html>
			<h1>Hello</h1>
			<p>現在時刻: %s</p>
			<p>あなたは %d番目の閲覧者です。</p>
			</html>`,
			nt.Format("2006/01/02 15:04:05.000"),
			counter)
		fmt.Fprintf(w, res)
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         addr + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	fmt.Printf("end!")
}
