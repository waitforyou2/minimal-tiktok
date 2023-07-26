package response

/*
Response 基础响应
*/
type Response struct {
	//状态代码
	StatusCode int32 `json:"status_code"`
	//状态信息
	StatusMsg string `json:"status_msg,omitempty"`
}
