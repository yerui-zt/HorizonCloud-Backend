package vm

import (
	"net/http"

	"HorizonX/common/result"

	"HorizonX/api/internal/logic/vm"
	"HorizonX/api/internal/svc"
)

func GetImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := vm.NewGetImageLogic(r.Context(), svcCtx)
		resp, err := l.GetImage()
		result.HttpResult(r, w, resp, err)
	}
}
