package main

import (
	"HNG_PROJECT_2/controllers"
	"fmt"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", " GET, POST, DELETE, PUT, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization ")

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/classify-number/", controllers.NumberClassificationController)
	handler := corsMiddleware(mux)

	port := ":8080"
	http.ListenAndServe(port, handler)
	fmt.Println("server is running on :8080")
}
