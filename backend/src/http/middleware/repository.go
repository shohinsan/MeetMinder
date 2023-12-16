package middleware

import (
	"context"
	"net/http"
	"shohinsan/MeetMinder/src/http/handlers"
)

type MiddlewareRepository struct {
	httpRepo *handlers.Repository
}

func NewMiddlewareRepository(repo *handlers.Repository) *MiddlewareRepository {
	return &MiddlewareRepository{
		httpRepo: repo,
	}
}

func (m *MiddlewareRepository) RequiresAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || len(authHeader) < 7 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenStr := authHeader[7:]

		claims, err := m.httpRepo.AuthTokenRepo.ParseToken(tokenStr)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), handlers.ContextKey{}, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
