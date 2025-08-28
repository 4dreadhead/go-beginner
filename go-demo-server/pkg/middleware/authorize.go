package middleware

import (
	"context"
	"goadvancedserver/configs"
	"goadvancedserver/pkg/jwthandler"
	"net/http"
	"strings"
)

type key string

const ContextEmailKey key = "ContextEmailKey"

func Authorize(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			sendUnauthorized(w)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, data := jwthandler.NewJWT(config.Auth.Secret).Parse(token)

		if !isValid {
			sendUnauthorized(w)
			return
		}

		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func sendUnauthorized(w http.ResponseWriter) {
	status := http.StatusUnauthorized
	w.WriteHeader(status)
	_, _ = w.Write([]byte(http.StatusText(status)))
}
