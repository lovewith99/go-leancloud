package goleancloud

import "testing"

func TestServiceConv(t *testing.T) {
	serviceName := "sys-upseries"

	cli := NewLeanClient(leanAppId, leanAppKey, leanMasterKey).NewServiceContext()

	// 获取服务号
	var convId string
	if r, err := cli.GetService(serviceName, 0, 1); true {
		t.Log("GetService: ", r)
		if err != nil {
			t.Error(err)
		}

		if r != nil && len(r.Results) > 0 {
			convId = r.Results[0].ObjectId
		}
	}

	if convId == "" {
		if r, err := cli.PostService(serviceName); true {
			t.Log("PostService: ", r)
			if err != nil {
				t.Error(err)
			} else {
				convId = r.ObjectId
			}

		}
	}

	t.Log("subscribe: ", cli.Subscribe(convId, "test01"))

	pubmsg := &ServiceConvBroadcastRequest{
		ConvId:     convId,
		FromClient: "system",
		Message:    "12345678",
	}

	if r, err := cli.Pub(pubmsg); true {
		t.Log("Pub: ", r)
		if err != nil {
			t.Error(err)
		}
	}

	t.Log("unsubscribe: ", cli.Unsubscribe(convId, "test01"))
}
