package tool

import "github.com/gin-gonic/gin"

// BadRequestResponse 400错误
func BadRequestResponse(c *gin.Context, msg ...string) {
	if len(msg) == 0 {
		msg = append(msg, "Bad Request")
	}

	c.JSON(400, gin.H{
		"code": 400,
		"msg":  msg,
	})
}

// InternalServerErrorResponse 500错误
func InternalServerErrorResponse(c *gin.Context, msg ...string) {
	if len(msg) == 0 {
		msg = append(msg, "Internal Server Error")
	}

	c.JSON(500, gin.H{
		"code": 500,
		"msg":  msg,
	})
}

// SuccessResponse 成功
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}
