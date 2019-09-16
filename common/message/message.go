package message

const (
	LoginMessgeType   = "LoginMes"
	LoginResponseType = "LoginRes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

type LoginMessage struct {
	UserID   int    `json:"user_id"`
	UserPwd  string `json:"user_pwd"`
	UserName string `json:"user_name"`
}

type LoginResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
