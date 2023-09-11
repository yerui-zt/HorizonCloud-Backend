// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	order "HorizonX/api/internal/handler/order"
	user "HorizonX/api/internal/handler/user"
	vm "HorizonX/api/internal/handler/vm"
	"HorizonX/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/logout",
					Handler: user.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Jwt.AccessSecret),
		rest.WithPrefix("/api/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/groups",
				Handler: vm.GetAllVMGroupsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/groups/:region",
				Handler: vm.GetVMGroupByRegionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/plans/:group_id",
				Handler: vm.GetVMPlanByGroupIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/images",
				Handler: vm.GetImageHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/vm"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/instance",
					Handler: vm.DeployVMInstanceHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Jwt.AccessSecret),
		rest.WithPrefix("/api/vm"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/:orderNo",
					Handler: order.GetOrderDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/:orderNo/paymentMethod",
					Handler: order.GetOrderPaymentMethodHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/:orderNo/pay",
					Handler: order.PayOrderHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Jwt.AccessSecret),
		rest.WithPrefix("/api/order"),
	)
}
