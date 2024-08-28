package internal

import (
	"net/http"
	"yutooaisdk/define"

	"github.com/gin-gonic/gin"
)

// type VersionParams struct {
// 	// 响应状态
// 	Status *define.PublicReply `json:"status"`
// 	// SDK版本
// 	SdkVersion string `json:"sdkVersion"`
// 	// 识别模型版本
// 	IdentifyVersion string `json:"identifyVersion"`
// 	// 匹配规则版本
// 	MatchingVersion string `json:"matchingVersion"`
// 	// 相机版本
// 	CameraVersion string `json:"cameraVersion"`
// 	// 机器码版本
// 	MachineVersion string `json:"machineVersion"`
// 	// 识别模型子版本（生鲜-fresh、零食-snack、通用-general）
// 	IdentifySubVersion string `json:"identifySubVersion"`
// }

// 1.该接口在初始化之前可调用
// 2.初始化之前调用该接口无法获取到识别算法的版本信息
func GetVersion(c *gin.Context) {
	//Todo
	c.JSON(http.StatusOK, gin.H{
		"status": &define.PublicReply{
			Code: 0,
			Msg:  "",
		},
		"sdkVersion":         "1.0.0.0",
		"identifyVersion":    "1.0.0.0",
		"matchingVersion":    "1.0.0.0",
		"cameraVersion":      "1.0.0.0",
		"machineVersion":     "1.0.0.0",
		"identifySubVersion": "snack",
	})
}
