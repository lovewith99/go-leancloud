package goleancloud

type PostServiceConvResponse struct {
	Code      int    `json:"code"`
	Error     string `json:"error"`
	ObjectId  string `json:"objectId"`
	CreatedAt string `json:"createdAt"`
}

type ServiceConvBroadcastRequest struct {
	ConvId     string `json:"-"`
	FromClient string `json:"from_client"`
	Message    string `json:"message"`
	Push       string `json:"push"`
}

type ServiceConvBroadcastResponse struct {
	Code      int    `json:"code"`
	Error     string `json:"error"`
	MsgId     string `json:"msg-id"`
	Timestamp int64  `json:"timestamp"`
}
