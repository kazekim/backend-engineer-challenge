package begincontext

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"net/http"
)

const (
	ErrHTTPCode = 599
)

type BaseResponse struct {
	Code    grerrors.ErrorCode `json:"code"`
	Message string             `json:"message,omitempty"`
	Data    interface{}        `json:"data"`
}

//wrapResponse create default response with code and message
func (c *Context) wrapResponse(v interface{}) BaseResponse {
	resp := BaseResponse{
		Code:    grerrors.SuccessCode,
		Message: "",
		Data:    v,
	}
	return resp
}

//CreateResponseSuccess create default response from struct
func (c *Context) CreateResponseSuccess(v interface{}) {
	resp := c.wrapResponse(v)
	c.JSON(http.StatusOK, resp)
}

func (c *Context) CreateResponseError(vErr grerrors.Error) {
	resp := BaseResponse{
		Code:    vErr.Code(),
		Message: vErr.Message(),
	}
	c.JSON(ErrHTTPCode, resp)
}
