package e

const (
	SYSTEM_ERROR = -1
	OK           = 0

	NOT_LOGIN  = 100
	NOT_AUTH   = 101
	LOGIN_FAIL = 102

	INVALID_PARAMS = 200
)

var MsgFlags = map[int]string{
	SYSTEM_ERROR:   "系统错误",
	OK:             "ok",
	NOT_LOGIN:      "未登录",
	NOT_AUTH:       "未授权",
	LOGIN_FAIL:     "登录失败",
	INVALID_PARAMS: "请求参数错误",
}

type ResponseData struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Data    any    `json:"data"`
}

func R(code int, data any) ResponseData {
	return ResponseData{
		ErrCode: code,
		ErrMsg:  MsgFlags[code],
		Data:    data,
	}
}
