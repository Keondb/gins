package res

import (
	"github.com/gin-gonic/gin"
	"gvb_server/utils"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}
type ListResponse[T any] struct {
	List  T     `json:"list"`
	Total int64 `json:"total"`
}

const (
	Success = 200
	Error   = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}
func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "操作成功", c)
}
func OkWithList(list any, total int64, c *gin.Context) {
	OkWithData(ListResponse[any]{
		List:  list,
		Total: total,
	}, c)
}
func OkWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}
func OkWith(c *gin.Context) {
	Result(Success, map[string]any{}, "操作成功", c)
}
func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}
func FailWithData(data any, c *gin.Context) {
	Result(Error, data, "操作成功", c)
}
func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}
func FailWithError(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, c)
}
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
	} else {
		Result(Error, map[string]any{}, "未知错误", c)
	}

}
