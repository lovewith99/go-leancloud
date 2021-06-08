package goleancloud

import "testing"

const (
	leanAppId     string = ""
	leanAppKey    string = ""
	leanMasterKey string = ""
)

func TestLeanCloudPush(t *testing.T) {
	cli := NewLeanClient(leanAppId, leanAppKey, leanMasterKey)
	pushBody := PushBody{
		Channels: []interface{}{"doll1000000052"},
		Prod:     "dev",
	}
	pushBody.SetData("aaa", "bbb", 2, map[string]interface{}{
		"action":   "com.youcompany.xxx",
		"push_key": "RRR",
	})
	cli.Push(&pushBody)
}
