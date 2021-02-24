package ginx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 1
	ERROR   = 0

	SuccessMsg = "ok"
	ErrorMsg   = "error"
)

// PageResult page result
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

// Option is response
type Option func(*Response)

// Response request rest response struct
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Gin struct {
	Ctx *gin.Context
}

// Response setting gin.JSON
func (g *Gin) Response(errCode int, errMsg string, data interface{}) {
	g.Ctx.JSON(http.StatusOK, Response{
		Code: errCode,
		Msg:  errMsg,
		Data: data,
	})
	return
}

func (g *Gin) Ok(data interface{}) {
	g.Response(SUCCESS, SuccessMsg, data)
	return
}

func (g *Gin) Error() {
	g.Response(ERROR, ErrorMsg, nil)
	return
}

// Ok set gin.JSON with success msg
func Ok(ctx *gin.Context, msg string) {
	JSON(ctx, WithMsg(msg))
}

// Success set gin.JSON with data
func Success(ctx *gin.Context, data interface{}) {
	JSON(ctx, WithMsg(SuccessMsg), WithData(data))
}

// Failure set gin.JSON with error msg
func Failure(ctx *gin.Context, msg string) {
	JSON(ctx, WithCode(ERROR), WithMsg(msg))
}

// JSON setting gin.JSON
func JSON(c *gin.Context, opts ...Option) {
	r := &Response{
		Code: SUCCESS,
	}

	for _, o := range opts {
		o(r)
	}

	c.JSON(http.StatusOK, r)
	return
}

// WithCode set code
func WithCode(code int) Option {
	return func(r *Response) {
		r.Code = code
	}
}

// WithMsg set msg
func WithMsg(msg string) Option {
	return func(r *Response) {
		r.Msg = msg
	}
}

// WithData set data
func WithData(data interface{}) Option {
	return func(r *Response) {
		r.Data = data
	}
}
