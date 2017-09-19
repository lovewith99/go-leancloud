package main

import (
	"github.com/fpagyu/go-leancloud"
)

const (
	leanAppId     string = ""
	leanAppKey    string = ""
	leanMasterKey string = ""
)

var LeanCloud *go_leancloud.LeanClient = go_leancloud.GetLeanClient(leanAppId, leanAppKey, leanMasterKey)

func main() {
	pushBody := go_leancloud.PushBody{}
	pushBody.Channels = []string{"doll1000000052"}
	pushBody.Prod = "dev"
	pushBody.SetData("aaa", "bbb", 2, map[string]interface{}{
		"action":   "com.youcompany.xxx",
		"push_key": "RRR",
	})
	LeanCloud.Push(&pushBody)
}
