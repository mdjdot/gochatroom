package common

const (
	// LoginMessage 登录信息
	LoginMessage = iota + 1
	// RegisterMessage 注册信息
	RegisterMessage
	// RequestMessage 请求信息
	RequestMessage
	// ResponseMessage 相应信息
	ResponseMessage

	// LoginSuccess 登录成功
	LoginSuccess = true
	// LoginFailed 登录失败
	LoginFailed = false
)

// Message .
type Message struct {
	Type int
	Data string
}

// LoginRespMessage .
type LoginRespMessage struct {
	Result bool
}

// RegisterRespMessage .
type RegisterRespMessage struct {
	Result bool
}
