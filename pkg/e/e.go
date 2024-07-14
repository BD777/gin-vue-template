package e

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RetCode int64

const (
	RetCodeSystem       RetCode = -1
	RetCodeOK           RetCode = 0
	RetCodeDataNotFound RetCode = 1

	RetCodeNotLogin      RetCode = 100
	RetCodeNotAuth       RetCode = 101
	RetCodeLoginFail     RetCode = 102
	RetCodeInvalidTenant RetCode = 103

	RetCodeInvalidParams RetCode = 200
	RetCodeNotSupport    RetCode = 201
	RetCodeDuplication   RetCode = 202
)

var retCodeNames = map[RetCode]string{
	RetCodeSystem:        "系统错误",
	RetCodeOK:            "ok",
	RetCodeDataNotFound:  "未找到数据",
	RetCodeNotLogin:      "未登录",
	RetCodeNotAuth:       "未授权",
	RetCodeLoginFail:     "登录失败",
	RetCodeInvalidTenant: "无效租户",
	RetCodeInvalidParams: "请求参数错误",
	RetCodeNotSupport:    "未支持",
	RetCodeDuplication:   "已有重复数据",
}

func (c RetCode) String() string {
	return retCodeNames[c]
}

type ResponseData struct {
	ErrCode RetCode `json:"errcode"`
	ErrMsg  string  `json:"errmsg"`
	Data    any     `json:"data"`
}

func R(code RetCode, data any) ResponseData {
	return ResponseData{
		ErrCode: code,
		ErrMsg:  code.String(),
		Data:    data,
	}
}

func OK(data any) ResponseData {
	return ResponseData{
		ErrCode: RetCodeOK,
		ErrMsg:  RetCodeOK.String(),
		Data:    data,
	}
}

func GinR(c *gin.Context, code RetCode, data any) {
	c.JSON(http.StatusOK, R(code, data))
}

func GinOK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, OK(data))
}
