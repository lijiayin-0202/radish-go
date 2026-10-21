package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 统一响应码，与前端 utils/request.js 约定保持一致
const (
	SUCCESS   = 0
	ERROR     = 500
	AUTHERROR = 401
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{Code: code, Data: data, Msg: msg})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]any{}, "操作成功", c)
}

func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, msg, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]any{}, "操作失败", c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(ERROR, map[string]any{}, msg, c)
}

func FailWithCode(code int, msg string, c *gin.Context) {
	Result(code, map[string]any{}, msg, c)
}
