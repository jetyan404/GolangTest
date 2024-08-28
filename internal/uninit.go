package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1.程序退出前务必调用一次该接口，否则部分学习数据可能丢失
func UnInit(c *gin.Context) {
	//Todo
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
	})
}
