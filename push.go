package goleancloud

import (
	"encoding/json"
)

// Document: https://leancloud.cn/docs/push_guide.html#推送消息
type PushBody struct {
	// required
	Data interface{} `json:"data"` // 推送的内容数据，JSON 对象

	// optional
	Channels           []interface{} `json:"channels,omitempty"`            // 推送给哪些频道，将作为条件加入 where 对象
	ExpirationInterval string        `json:"expiration_interval,omitempty"` // 消息过期的相对时间，从调用 API 的时间开始算起，单位是秒
	ExpirationTime     string        `json:"expiration_time,omitempty"`     // 消息过期的绝对日期时间
	NotificationId     string        `json:"notification_id,omitempty"`     // 自定义推送 id，最长 16 个字符且只能由英文字母和数字组成
	Prod               string        `json:"prod,omitempty"`                // 仅对iOS有效, 设置使用开发证书, 还是生产证书
	PushTime           string        `json:"push_time,omitempty"`           // 定期推送时间
	ReqId              string        `json:"req_id,omitempty"`              // 自定义请求 id，最长 16 个字符且只能由英文字母和数字组成
	Where              interface{}   `json:"where,omitempty"`               // 检索 _Installation 表使用的查询条件，JSON 对象。
}

func (p *PushBody) Buffer() []byte {
	r, err := json.Marshal(&p)
	if err != nil {
		return nil
	}
	return r
}

func (p *PushBody) SetData(title, msg string, device_type int, d map[string]interface{}) {
	// device_type: (1:ios; 2:android; 3:wp; 4:ios and android; 0: ios and android and wp)
	iosIns := IosPushData{}
	andIns := AndroidPushData{}
	wpIns := WpPushData{}

	switch device_type {
	case 1:
		iosIns.Object(title, msg, d)
		p.Data = iosIns
	case 2:
		andIns.Object(title, msg, d)
		p.Data = andIns
	case 3:
		wpIns.Object(title, msg, d)
		p.Data = wpIns
	case 4:
		data := make(map[string]interface{})
		iosIns.Object(title, msg, d)
		andIns.Object(title, msg, d)
		data["ios"], data["android"] = iosIns, andIns
		p.Data = data
	case 0:
		data := make(map[string]interface{})
		iosIns.Object(title, msg, d)
		andIns.Object(title, msg, d)
		wpIns.Object(title, msg, d)
		data["ios"], data["android"], data["wp"] = iosIns, andIns, wpIns
		p.Data = data
	default:
		d["alert"] = msg
		d["title"] = title
		p.Data = d
	}
}

// Document: https://leancloud.cn/docs/push_guide.html#消息内容_Data
type PushData interface {
	PushType() string
}

type IosPushDataBase struct {
	Alter            interface{} `json:"alert"`                       // 消息内容,字符串; 如果alter本地化推送, 将alert参数从string替换为一个由本地化消息推送属性组成的json
	Category         string      `json:"category,omitempty"`          // 通知类型
	ThreadId         string      `json:"thread-id,omitempty"`         // 通知分类名称
	Badge            interface{} `json:"badge"`                       // 数字类型，未读消息数目，应用图标边上的小红点数字，可以是数字，也可以是字符串 "Increment"（大小写敏感）,
	Sound            string      `json:"sound,omitempty"`             // 声音文件名，前提在应用里存在
	ContentAvailable int         `json:"content-available,omitempty"` // 数字类型，如果使用 Newsstand，设置为 1 来开始一次后台下载
	MutableContent   int         `json:"mutable-content,omitempty"`   // 数字类型，用于支持 UNNotificationServiceExtension 功能，设置为 1 时启用
}

// alter 本地化推送
type IosPushDataAlter struct {
	Title          string   `json:"title"` // 标题
	TitleLocKey    string   `json:"title-loc-key,omitempty"`
	SubTitle       string   `json:"sub-title"` // 副标题
	SubTitleLocKey string   `json:"sub-title-loc-key,omitempty"`
	Body           string   `json:"body"` // 消息内容
	ActionLocKey   string   `json:"action-loc-key,omitempty"`
	LocKey         string   `json:"loc-key,omitempty"`
	LocArgs        []string `json:"loc-args,omitempty"`
	LaunchImage    string   `json:"launch-image"`
}

// 苹果官方文档方式构建推送参数(暂不支持)
//type ApsPushData struct {
//	Aps   IosPushDataBase `json:"aps"`
//	Other interface{}     `json:"other,omitempty"` // 自定义字段
//}
//
//func (p ApsPushData) PushType() string {
//	return "ios"
//}

type IosPushData struct {
	IosPushDataBase
	PushKey interface{} `json:"PushKey,omitempty"`
	Data    interface{} `json:"data,omitempty"` // 自定义字段
}

func (p IosPushData) PushType() string {
	return "ios"
}

func (p *IosPushData) Object(title, msg string, d map[string]interface{}) {
	p.Alter = msg
	p.Badge = "Increment"

	if d == nil {
		return
	}

	if v, ok := d["badge"]; ok {
		p.Badge = v
	} else {
		p.Badge = "Increment"
	}

	if v, ok := d["category"]; ok {
		category, _ := v.(string)
		p.Category = category
	}

	if v, ok := d["thread-id"]; ok {
		threadId, _ := v.(string)
		p.ThreadId = threadId
	}

	if v, ok := d["sound"]; ok {
		sound, _ := v.(string)
		p.Sound = sound
	}

	if v, ok := d["PushKey"]; ok {
		pushKey, _ := v.(string)
		p.PushKey = pushKey
	}

	if v, ok := d["content-available"]; ok {
		contentAvailable, _ := v.(int)
		p.ContentAvailable = contentAvailable
	}

	if v, ok := d["mutable-content"]; ok {
		mutableContent, _ := v.(int)
		p.MutableContent = mutableContent
	}

	if v, ok := d["data"]; ok {
		p.Data = v
	}

	return
}

type AndroidPushData struct {
	Alter   string      `json:"alert"`             // 消息内容
	Title   string      `json:"title"`             // 显示在通知栏标题
	Action  string      `json:"action"`            // com.your-company.push
	Silent  bool        `json:"silent,omitempty"`  // 用于控制是否关闭通知栏提醒, 默认为false,即不关闭通知栏提醒
	PushKey string      `json:"PushKey,omitempty"` // 自定义字段
	Data    interface{} `json:"data,omitempty"`    // 自定义字段
}

func (p AndroidPushData) PushType() string {
	return "android"
}

func (p *AndroidPushData) Object(title, msg string, d map[string]interface{}) {
	p.Alter = msg
	p.Title = title

	if d == nil {
		return
	}

	if v, ok := d["action"]; ok {
		action, _ := v.(string)
		p.Action = action
	}

	if v, ok := d["silent"]; ok {
		silent, _ := v.(bool)
		p.Silent = silent
	}

	if v, ok := d["PushKey"]; ok {
		pushKey, _ := v.(string)
		p.PushKey = pushKey
	}

	if v, ok := d["data"]; ok {
		p.Data = v
	}

	return
}

type WpPushData struct {
	Alter   string `json:"alert"`    // 消息内容
	Title   string `json:"title"`    // 显示在通知栏标题
	WpParam string `json:"wp-param"` // "/chat.xaml?NavigatedFrom=Toast Notification"
}

func (p WpPushData) PushType() string {
	return "wp"
}

func (p *WpPushData) Object(title, msg string, d map[string]interface{}) {
	p.Alter = msg
	p.Title = title

	if d == nil {
		return
	}

	if v, ok := d["wp-param"]; ok {
		wpParam, _ := v.(string)
		p.WpParam = wpParam
	}

	return
}
