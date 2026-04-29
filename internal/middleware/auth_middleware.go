package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/sriraghariharan/gotasks/internal/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type contextKey string

const UserIDKey contextKey = "userId"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Expect: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid Authorization format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		// verify token
		userID, err := utils.VerifyJwt(tokenString)
		if err != nil {
			http.Error(w, "invalid or expired token", http.StatusUnauthorized)
			return
		}

		// convert userId to object id
		userObjectID, err := bson.ObjectIDFromHex(userID)
		if err != nil {
			http.Error(w, "invalid User", http.StatusUnauthorized)
			return
		}

		// attach userId to request context
		ctx := context.WithValue(r.Context(), UserIDKey, userObjectID)

		// pass to next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}