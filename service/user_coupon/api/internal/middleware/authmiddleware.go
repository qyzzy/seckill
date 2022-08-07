package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	logx.Info("auth middle")
	return func(w http.ResponseWriter, r *http.Request) {
		// Passthrough to next handler if need
		next(w, r)
	}
}
