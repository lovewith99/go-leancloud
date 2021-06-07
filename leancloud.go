package go_leancloud

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

type LeanClient struct {
	AppId     string
	AppKey    string
	MasterKey string
	*http.Client
}

func NewLeanClient(appId, appKey, masterKey string) *LeanClient {
	return &LeanClient{
		AppId:     appId,
		AppKey:    appKey,
		MasterKey: masterKey,
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
	req.Header.Set("X-LC-Key", cli.AppKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (cli *LeanClient) Push(p *PushBody) error {
	req, err := cli.NewRequest("POST", "https://api.leancloud.cn/1.1/push", p.Buffer())
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

// // PostServiceConv 创建服务号
// func (cli *LeanClient) PostServiceConv(name string) error {
// 	body := map[string]interface{}{
// 		"name": name,
// 	}
// 	buf, _ := json.Marshal(body)
// 	req, err := cli.NewRequest("POST", "https://api.leancloud.cn/1.2/rtm/service-conversations", buf)
// 	if err != nil {
// 		return err
// 	}
// 	resp, err := cli.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	// resp, err = ioutil.ReadAll(resp.Body)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	return nil
// 	// buf := json.Marshal
// }
