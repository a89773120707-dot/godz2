package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		n := rand.IntN(6) + 1
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%d", n)
	})
	fmt.Println("Сервер запущен на порту: 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Printf("Ошибка сервера: %v", err)
		return
	}
}
