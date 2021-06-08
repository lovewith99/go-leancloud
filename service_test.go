package goleancloud

import "testing"

func TestServiceConv(t *testing.T) {
	serviceName := "sys-upseries"

	cli := NewLeanClient(leanAppId, leanAppKey, leanMasterKey).NewServiceContext()
	r2, err := cli.GetServiceConv(serviceName, 0, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log("response: ", r2)

	var convId string
	if len(r2.Results) == 0 {
		r, err := cli.PostServiceConv(serviceName)
		if err != nil {
			t.Error(err)
		}
		t.Log("response: ", r)
		convId = r.ObjectId
	} else {
		convId = r2.Results[0].ObjectId
	}

	r1, err := cli.BroadcastServiceConv(&ServiceConvBroadcastRequest{
		ConvId:     convId,
		FromClient: "system",
		Message:    "12345678",
	})
	if err != nil {
		t.Error(err)
	}

	t.Log("response1: ", r1)
}
