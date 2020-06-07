package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Rota OK"))
	})

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		data := map[string]string{
			"message": "Trabalhando com objeto json",
		}

		result, _ := json.Marshal(data)

		w.Write([]byte(result))
	})

	http.ListenAndServe(":3001", nil)
}