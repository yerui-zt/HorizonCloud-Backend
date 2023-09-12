package xerr

//grpc 的 err 对应的错误码其实就是一个 uint32 ， 我们自己定义错误用 uint32 然后在 rpc 的全局拦截器返回时候转成 grpc 的 err

const OK string = "OK"

// 通用错误码
const (
	SERVER_COMMON_ERROR string = "SERVER_COMMON_ERROR"
	REQUEST_PARAM_ERROR string = "REQUEST_PARAM_ERROR"
	UNAUTHORIZED        string = "UNAUTHORIZED"
	DB_ERROR            string = "DB_ERROR"
	DB_NOT_FOUND_ERROR  string = "DB_NOT_FOUND_ERROR"

	JWT_ISSUE_ERROR string = "JWT_ISSUE_ERROR"
)

// RPC模块细分错误码，如用户模块，订单模块等
// 用户模块
const (
	USER_PASSWORD_ERROR  string = "USER_PASSWORD_ERROR"
	USER_EXIST_ERROR     string = "USER_EXIST_ERROR"
	USER_NOT_FOUND_ERROR string = "USER_NOT_FOUND_ERROR"
)

// 订单模块
const (
	ORDER_CREATE_ERROR       string = "ORDER_CREATE_ERROR"
	ORDER_PLAN_NOT_FOUND            = "PLAN_NOT_FOUND"
	ORDER_VM_GROUP_NOT_FOUND        = "VM_GROUP_NOT_FOUND"
	ORDER_VM_IMAGE_NOT_FOUND        = "IMAGE_NOT_FOUND"
	ORDER_NOT_FOUND                 = "ORDER_NOT_FOUND"
	ORDER_HAS_PAID                  = "ORDER_HAS_BEEN_PAID"
)

// 支付模块
const (
	PAYMENT_METHOD_NOT_FOUND string = "PAYMENT_METHOD_NOT_FOUND"
	PAYMENT_CREATE_ERROR     string = "PAYMENT_CREATE_ERROR"
)
