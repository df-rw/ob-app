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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := flag.Int("p", 8080, "webserver port")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/now", now)

	corsMux := corsMiddleware(mux)

	fmt.Println("Listening on port", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), corsMux))
}
