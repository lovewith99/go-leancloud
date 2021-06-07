package goleancloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type LeanClient struct {
	AppId     string
	AppKey    string
	MasterKey string
	Endpoint  string
	*http.Client
}

func NewLeanClient(appId, appKey, masterKey string) *LeanClient {
	return &LeanClient{
		AppId:     appId,
		AppKey:    appKey,
		MasterKey: masterKey,
		Endpoint:  "https://api.leancloud.cn",
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func GetLeanClient(appId, appKey, masterKey string) *LeanClient {
	return NewLeanClient(appId, appKey, masterKey)
}

func (cli *LeanClient) NewRequest(method, apiUrl string, buf []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, apiUrl, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-LC-Id", cli.AppId)
	req.Header.Set("X-LC-Key", cli.MasterKey+",master")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (cli *LeanClient) Push(p *PushBody) error {
	req, err := cli.NewRequest("POST", cli.Endpoint+"/1.1/push", p.Buffer())
	if err != nil {
		return err
	}

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// PostServiceConv 创建服务号
func (cli *LeanClient) PostServiceConv(name string) (*PostServiceConvResponse, error) {
	body := map[string]interface{}{
		"name": name,
	}
	buf, _ := json.Marshal(body)
	req, err := cli.NewRequest("POST", cli.Endpoint+"/1.2/rtm/service-conversations", buf)
	if err != nil {
		return nil, err
	}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r PostServiceConvResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	if r.Error != "" {
		return &r, errors.New(r.Error)
	}

	return &r, nil
}

// BroadcastServiceConv 给所有订阅者发消息
func (cli *LeanClient) BroadcastServiceConv(request *ServiceConvBroadcastRequest) (*ServiceConvBroadcastResponse, error) {
	buf, _ := json.Marshal(request)
	req, err := cli.NewRequest("POST",
		fmt.Sprintf("%s/1.2/rtm/service-conversations/%s/broadcasts",
			cli.Endpoint, request.ConvId), buf)
	if err != nil {
		return nil, err
	}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r ServiceConvBroadcastResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	if r.Error != "" {
		return &r, errors.New(r.Error)
	}

	return &r, nil
}
