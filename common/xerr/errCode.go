package xerr

var code map[string]uint32

func init() {
	code = make(map[string]uint32)
	code[OK] = 200
	code[SERVER_COMMON_ERROR] = 400
	code[REQUEST_PARAM_ERROR] = 400
	code[UNAUTHORIZED] = 401
	code[GENERAL_NOT_FOUND_ERROR] = 404
	code[DB_ERROR] = 500
	code[DB_NOT_FOUND_ERROR] = 500
	code[JWT_ISSUE_ERROR] = 500
	code[MQ_PUBLISH_ERROR] = 500
	code[JSON_UNMARSHAL_ERROR] = 500

	// 用户模块
	code[USER_PASSWORD_ERROR] = 400
	code[USER_EXIST_ERROR] = 400
	code[USER_NOT_FOUND_ERROR] = 400

	// 订单模块
	code[ORDER_CREATE_ERROR] = 500
	code[ORDER_NOT_FOUND] = 400
	code[ORDER_PLAN_NOT_FOUND] = 400
	code[ORDER_VM_GROUP_NOT_FOUND] = 400
	code[ORDER_VM_IMAGE_NOT_FOUND] = 400
	code[ORDER_HAS_PAID] = 400
	code[ORDER_STATUS_ERROR] = 400
	code[ORDER_ITEM_PARSE_ERROR] = 500

	// 支付模块
	code[PAYMENT_METHOD_NOT_FOUND] = 400
	code[PAYMENT_CREATE_ERROR] = 500

	// VM 模块
	code[HYPERVISOR_NODE_UNAVAILABLE] = 400
	code[OS_IMAGE_NOT_AVAILABLE] = 400
	code[PROXMOX_VM_CREATE_FAILED] = 500
	code[PROXMOX_VM_FETCH_ERROR] = 500
	code[PROXMOX_VM_CONFIG_ERROR] = 500
	code[IP_NO_AVAILABLE_ADDR_ERROR] = 500

	// SSH_KEY 模块
	code[INVALID_SSH_KEY] = 400
}

// MapErrMsg 根据错误类型返回错误码
func Code(errType string) uint32 {
	if c, ok := code[errType]; ok {
		return c
	}
	return 400
}

// IsCodeErr 判断是否是已定义的错误码
func IsCodeErr(errType string) bool {
	if _, ok := code[errType]; ok {
		return true
	} else {
		return false
	}
}
