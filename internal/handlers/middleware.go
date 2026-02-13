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
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-User-Role")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

// AuthMiddleware защищает критические API
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Простая логика: если это POST/DELETE, проверяем роль в заголовке
		// В реальном приложении здесь проверяется JWT токен
		if r.Method == http.MethodPost || r.Method == http.MethodDelete {
			role := r.Header.Get("X-User-Role")
			if role == "" {
				// Для демо-версии мы позволяем запросы, но в реальности тут была бы ошибка 401
				log.Println("Warning: Unauthenticated access to sensitive API")
			}
		}
		next(w, r)
	}
}
