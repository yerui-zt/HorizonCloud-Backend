package xerr

var code map[string]uint32

func init() {
	code = make(map[string]uint32)
	code[OK] = 200
	code[SERVER_COMMON_ERROR] = 400
	code[REQUEST_PARAM_ERROR] = 400
	code[UNAUTHORIZED] = 401
	code[DB_ERROR] = 500
	code[DB_NOT_FOUND_ERROR] = 500
	code[JWT_ISSUE_ERROR] = 500

	// 用户模块
	code[USER_PASSWORD_ERROR] = 400
	code[USER_EXIST_ERROR] = 400
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