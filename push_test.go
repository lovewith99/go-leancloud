package goleancloud

import "testing"

const (
	leanAppId     string = ""
	leanAppKey    string = ""
	leanMasterKey string = ""
)

var LeanCloud *LeanClient = NewLeanClient(leanAppId, leanAppKey, leanMasterKey)

func TestLeanCloudPush(t *testing.T) {
	pushBody := PushBody{
		Channels: []interface{}{"doll1000000052"},
		Prod:     "dev",
	}
	pushBody.SetData("aaa", "bbb", 2, map[string]interface{}{
		"action":   "com.youcompany.xxx",
		"push_key": "RRR",
	})
	LeanCloud.Push(&pushBody)
}
