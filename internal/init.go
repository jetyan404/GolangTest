package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1.建议返回状态成功后不再调用本接口
func Init(c *gin.Context) {
	//Todo
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
	})
}
