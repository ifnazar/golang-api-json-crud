package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

// HealthStruct - structure to represent application status
type HealthStruct struct {
	Status string `json:"status"`
}

func rootEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func health(w http.ResponseWriter, r *http.Request) {
	var health = HealthStruct{Status: "Ok"}
	json.NewEncoder(w).Encode(health)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(GetIP(r) + " >> " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func handleRequests() {
	// initialize router
	r := mux.NewRouter()

	// add middlewares
	r.Use(loggingMiddleware)

	// endpoints
	r.HandleFunc("/", rootEndPoint)
	r.HandleFunc("/health", health)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", r)
	log.Fatal(err)
}

func main() {
	handleRequests()
}
