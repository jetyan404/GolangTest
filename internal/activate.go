package internal

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActivateModel struct {
	ActivateInfo *ActivateInfo `json:"activateInfo"`
}

// ActivateInfo 激活信息
type ActivateInfo struct {
	// 激活码【必填】
	CDKey string `json:"cdKey"`
	// 门店编码【可选】
	ShopCode string        `json:"shopCode,omitempty"`
	ShopInfo *ShopBaseInfo `json:"shopInfo,omitempty"`
}

// ShopBaseInfo 门店基础信息
type ShopBaseInfo struct {
	// 门店联系人【可选】
	Contact string `json:"contact,omitempty"`
	// 门店名称【可选】
	Name string `json:"name,omitempty"`
	// 联系人电话【可选】
	Phone string `json:"phone,omitempty"`
}

// 1.建议仅在Init接口返回未激活状态码时调用该接口
func Activate(c *gin.Context) {
	//Todo
	var data ActivateModel
	var code int32
	boby, _ := io.ReadAll(c.Request.Body)
	err := json.Unmarshal(boby, &data)
	if err != nil {
		code = -1
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  data,
	})
}
