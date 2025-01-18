package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/douyin/app/favorite/common/errno"
)

type Response struct {
	Code int32
	Msg  string
	Data any
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code: errno.Success.Code,
		Msg:  errno.Success.Msg,
		Data: data,
	})
}

func Error(c *gin.Context, err error) {
	e := errno.Errno{}
	if errors.As(err, &e) {
		c.JSON(http.StatusOK, Response{
			Code: e.Code,
			Msg:  e.Msg,
		})
	}

	resp := errno.InternalServer.WithMessage(err.Error())
	c.JSON(http.StatusOK, Response{
		Code: resp.Code,
		Msg:  resp.Msg,
	})
}
