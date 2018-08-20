package e

var MsgFlags = map[int]string{
	SUCCESS:              "ok",
	INNER_ERROR:          "inner error",
	ERROR_INVALID_PARAMS: "请求参数错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[INNER_ERROR]
}
