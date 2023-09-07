package user

import (
	"net/http"

	"HorizonX/common/result"

	"HorizonX/api/internal/logic/user"
	"HorizonX/api/internal/svc"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout()
		result.HttpResult(r, w, resp, err)
	}
}
