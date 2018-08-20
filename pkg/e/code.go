package e

// 0 成功
// 3位数 内部错误
// 5位数 外部错误
const (
	SUCCESS = 0

	INNER_ERROR = 100

	ERROR_INVALID_PARAMS = 10001
)

var (
	WX_CODE_EMPTY   = New(ERROR_INVALID_PARAMS, "WX Code Empty")
	WX_OPENID_EMPTY = New(ERROR_INVALID_PARAMS, "WX Openid Empty")
)
