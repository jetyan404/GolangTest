package define

// YoYoApiPath 由图API接口地址
const (
	// GetVersion 获取版本信息
	GetVersion string = "/grpc/version"
	// Init 服务初始化接口
	Init string = "/grpc/init"
	// UnInit 服务释放接口
	UnInit string = "/grpc/uninit"
	// Activate 设备激活
	Activate string = "/grpc/activate"
)

// PublicReply 公共响应（出参）
type PublicReply struct {
	// 状态码
	Code int32 `json:"code"`
	// 状态描述
	Msg string `json:"msg"`
}
