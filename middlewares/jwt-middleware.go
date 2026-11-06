package auth

import (
	"context"
	"encoding/json"
	"github/giovanifz/todolist-go/controllers"
	"net/http"
	"strings"
)

type contextKey string

const ClaimsContextKey contextKey = "claims"

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			answerError(w, "Authentication missing", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
			answerError(w, "Invalid format. Use: Bearer <token>", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		claims, err := controllers.ValidateAccessToken(tokenString)
		if err != nil {
			answerError(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func answerError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"erro":   message,
		"codigo": statusCode,
	})
}
