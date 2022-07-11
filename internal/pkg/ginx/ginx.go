package ginx

import (
	"fmt"
	"go-server/pkg/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func result(c *gin.Context, code int, msg string, data interface{}) {
	if data == nil {
		data = struct{}{}
	}

	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Success(c *gin.Context, data interface{}) {
	result(c, 0, "", data)
}

func Fail(c *gin.Context, code int, msg string) {
	result(c, code, msg, nil)
}

func ResError(c *gin.Context, err error) {
	if e, ok := err.(*errors.InnerError); ok {
		Fail(c, e.Code, e.Message)
	} else {
		logrus.Errorf("%+v", err)
		Fail(c, 500, "Internal Error")
	}
}

// ResList Response data with list object
func ResList(c *gin.Context, v interface{}) {
	Success(c, ListResult{List: v})
}

// ResPage Response pagination data object
func ResPage(c *gin.Context, v interface{}, pr *PaginationResult) {
	list := ListResult{
		List:       v,
		Pagination: pr,
	}
	Success(c, list)
}

// ParseParamID Param returns the value of the URL param
func ParseParamID(c *gin.Context, key string) uint64 {
	val := c.Param(key)
	id, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0
	}
	return id
}

// ParseJSON Parse body json data to struct
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return Wrap400ResponseError(fmt.Sprintf("Parse request json failed: %s", err.Error()))
	}
	return nil
}

// ParseQuery Parse query parameter to struct
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return Wrap400ResponseError(fmt.Sprintf("Parse request query failed: %s", err.Error()))
	}
	return nil
}

// ParseForm Parse body form data to struct
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return Wrap400ResponseError(fmt.Sprintf("Parse request form failed: %s", err.Error()))
	}
	return nil
}
