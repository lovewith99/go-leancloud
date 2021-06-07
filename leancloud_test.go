package goleancloud

import "testing"

const (
	leanAppId     = ""
	leanAppKey    = ""
	leanMasterKey = ""
	prod          = "dev"
)

func TestPostServiceConv(t *testing.T) {
	cli := NewLeanClient(leanAppId, leanAppKey, leanMasterKey)
	r, err := cli.PostServiceConv("test")
	if err != nil {
		t.Error(err)
	}

	t.Log("response: ", r)

	r1, err := cli.BroadcastServiceConv(&ServiceConvBroadcastRequest{
		ConvId:     "test",
		FromClient: "system",
		Message:    "12345678",
	})
	if err != nil {
		t.Error(err)
	}

	t.Log("response1: ", r1)
}
