package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", myMiddleware(myhandler))

	http.ListenAndServe(":8080", nil)

}

func myhandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello from Hendler")
	w.WriteHeader(http.StatusOK)
}

func myMiddleware(nex http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello from Middleware")
		nex(w, r)
	}

}
