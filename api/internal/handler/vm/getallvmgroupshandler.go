package vm

import (
	"net/http"

	"HorizonX/common/result"

	"HorizonX/api/internal/logic/vm"
	"HorizonX/api/internal/svc"
)

func GetAllVMGroupsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := vm.NewGetAllVMGroupsLogic(r.Context(), svcCtx)
		resp, err := l.GetAllVMGroups()
		result.HttpResult(r, w, resp, err)
	}
}
