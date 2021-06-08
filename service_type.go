package goleancloud

// leancloud 服务号
// doc: https://leancloud.cn/docs/realtime_rest_api_v2.html#hash-114070650

type PostServiceConvResponse struct {
	Code      int    `json:"code"`
	Error     string `json:"error"`
	ObjectId  string `json:"objectId"`
	CreatedAt string `json:"createdAt"`
}

type GetServiceConvResponse struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Results []struct {
		Name      string `json:"name"`
		Sys       bool   `json:"sys"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
		ObjectId  string `json:"objectId"`
	} `json:"results"`
}

type ServiceConvBroadcastRequest struct {
	ConvId     string `json:"-"`
	FromClient string `json:"from_client"`
	Message    string `json:"message"`
	Push       string `json:"push,omitempty"`
}

type ServiceConvBroadcastResponse struct {
	Code   int    `json:"code"`
	Error  string `json:"error"`
	Result struct {
		MsgId     string `json:"msg-id"`
		Timestamp int64  `json:"timestamp"`
	} `json:"result"`
}
