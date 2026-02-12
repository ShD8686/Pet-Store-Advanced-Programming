package handlers

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s - %s", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
		next(w, r)
	}
}

func CommonHeadersMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Заглушка авторизации
		token := r.Header.Get("Authorization")
		if token == "" && r.Method != http.MethodGet {
			// На реальном проекте здесь будет проверка JWT
			// http.Error(w, "Unauthorized", http.StatusUnauthorized)
			// return
		}
		next(w, r)
	}
}
