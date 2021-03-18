package response

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"vvvstore/internal/pkg/errno"
)


// 响应结构
type Response struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	RequestId string      `json:"request_id"`
	Timestamp int64       `json:"timestamp"`
}

func (r Response) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(b)
}

// 响应成功
func Success(ctx *gin.Context, msg string, data interface{}) {
	if len(msg) == 0 {
		msg = errno.OK.Message
	}
	resp := &Response{
		Code:      errno.OK.Code,
		Msg:       msg,
		Data:      data,
		RequestId: ctx.GetString("X-Request-Id"),
		Timestamp: time.Now().Unix(),
	}
	ctx.JSON(http.StatusOK, resp)
}

// 响应失败
func Fail(ctx *gin.Context, err error) {
	code, message := errno.Decode(err)
	resp := Response{
		Code:      code,
		Msg:       message,
		Data:      nil,
		RequestId: ctx.GetString("X-Request-Id"),
		Timestamp: time.Now().Unix(),
	}
	ctx.JSON(http.StatusOK, resp)
}
