package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var counter = 0

func main() {
	usageMsg := `Usage:
	<command> <address> <port>
`
	if len(os.Args) != 3 {
		fmt.Println(usageMsg)
		os.Exit(1)
	}

	addr := os.Args[1]
	port := os.Args[2]

	if addr == "" || port == "" {
		fmt.Println(usageMsg)
		os.Exit(1)
	}

	fmt.Println("start!")
	myPrivateIps := myPrivateIps()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		nt := time.Now()
		counter++
		res := fmt.Sprintf(
			`<html>
	<head>
		<style type="text/css">
		p {
			color: #1a1a1a;
		}
		</style>
	</head>
	<h1>Hello</h1>
	<p>現在時刻: %s</p>
	<p>あなたは %d番目の閲覧者です。</p>
	<div>from %s</div>
</html>`,
			nt.Format("2006/01/02 15:04:05.000"),
			counter,
			myPrivateIps)
		fmt.Fprint(w, res)
	})
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		fmt.Fprintf(w, "OK")
	})
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		nt := time.Now()
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello", "time": nt.Format("2006/01/02 15:04:05.000"), "IP": fmt.Sprintf("%s", myPrivateIps)})
	})
	r.HandleFunc("/greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		reqPathVars := mux.Vars(r)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello, " + reqPathVars["name"]})
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

// プライベートIP（一番最初のもの）を返す
func myPrivateIps() []string {
	netInterfaceAddresses, _ := net.InterfaceAddrs()

	ips := []string{}
	for _, netInterfaceAddress := range netInterfaceAddresses {
		networkIp, ok := netInterfaceAddress.(*net.IPNet)
		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
			ips = append(ips, networkIp.IP.String())
		}
	}
	return ips
}
