package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func now(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	w.Header().Set("Content-type", "text/html")
	fmt.Fprintf(w, "%s", t)
}

func then(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	t = t.Add(10 * time.Minute)

	w.Header().Set("Content-type", "text/html")
	fmt.Fprintf(w, "%s", t)
}

func main() {
	port := flag.Int("p", 8082, "webserver port")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/now", now)
	mux.HandleFunc("/api/then", then)
	mux.Handle("/", http.FileServer(http.Dir("./dist")))

	fmt.Println("Listening on port", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}
