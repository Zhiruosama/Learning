package main

import (
	"fmt"
	"log"
	"net/http"
)

// import (
// 	"log"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/ping", pong)
// 	log.Println("Starting http server ...")
// 	log.Fatal(http.ListenAndServe(":50052", nil))
// }

// func pong(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("pong"))
// }

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-type", "text/plain; charest=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	port := "8080"
	log.Printf("Starting RESTful API server on port %s...", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
