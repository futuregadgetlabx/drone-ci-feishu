package request

type GetTokenReq struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type GetTokenResp struct {
	Code              int64  `json:"code"`
	Expire            int64  `json:"expire"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
}

type SendMsgReq struct {
	ReceiveID string `json:"receive_id"`
	MsgType   string `json:"msg_type"`
	Content   string `json:"content"`
	UUID      string `json:"uuid"`
}
