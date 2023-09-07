package middleware

import (
	"HorizonX/api/internal/config"
	"HorizonX/rpc/identity/identityservice"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Config      config.Config
	IdentityRPC identityservice.IdentityService
}

func NewAuthMiddleware(c config.Config, identityRPC identityservice.IdentityService) *AuthMiddleware {
	return &AuthMiddleware{
		Config:      c,
		IdentityRPC: identityRPC,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Authorization := r.Header.Get("Authorization")
		// 获取 Bearer token
		token := strings.Split(Authorization, " ")[1]

		rpcResp, err := m.IdentityRPC.VerifyJWT(r.Context(), &identityservice.VerifyJWTReq{
			Token: token,
		})
		if err != nil {
			logx.Errorf("rpc verify jwt failed: %v", err)
			http.Error(w, "JWT AuthMiddleware Error", http.StatusInternalServerError)
			return
		}
		if rpcResp.Valid == false {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		reqCtx := r.Context()
		ctx := context.WithValue(reqCtx, "jwtToken", token)
		newReq := r.WithContext(ctx)
		next(w, newReq)
	}
}
