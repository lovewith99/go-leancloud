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
	ConvId     string   `json:"-"`
	FromClient string   `json:"from_client"`
	ToClients  []string `json:"to_clients,omitempty"`
	Message    string   `json:"message"`

	Transient bool        `json:"transient,omitempty"` // 是否为暂态消息，默认 false
	NoSync    bool        `json:"no_sync,omitempty"`   // 默认情况下消息会被同步给在线的 from_client 用户的客户端，设置为 true 禁用此功能
	PushData  interface{} `json:"push_data,omitempty"` // 以消息附件方式设置本条消息的离线推送通知内容。如果目标接收者使用的是 iOS 设备并且当前不在线，我们会按照该参数填写的内容来发离线推送。
	Priority  string      `json:"priority,omitempty"`  // 定义消息优先级，可选值为 high、normal、low，分别对应高、中、低三种优先级。该参数大小写不敏感，默认为高优先级 high。本参数仅对暂态消息或聊天室的消息有效，高优先级下在服务端与用户设备的连接拥塞时依然排队。
	// Push      string      `json:"push,omitempty"`
}

type ServiceConvBroadcastResponse struct {
	Code   int    `json:"code"`
	Error  string `json:"error"`
	Result struct {
		MsgId     string `json:"msg-id"`
		Timestamp int64  `json:"timestamp"`
	} `json:"result"`
}
