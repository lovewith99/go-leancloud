package main

import (
	goleancloud "github.com/fpagyu/go-leancloud"
)

const (
	leanAppId     string = ""
	leanAppKey    string = ""
	leanMasterKey string = ""
)

var LeanCloud *goleancloud.LeanClient = goleancloud.GetLeanClient(leanAppId, leanAppKey, leanMasterKey)

func main() {
	pushBody := goleancloud.PushBody{}
	pushBody.Channels = []interface{}{"doll1000000052"}
	pushBody.Prod = "dev"
	pushBody.SetData("aaa", "bbb", 2, map[string]interface{}{
		"action":   "com.youcompany.xxx",
		"push_key": "RRR",
	})
	LeanCloud.Push(&pushBody)
}
