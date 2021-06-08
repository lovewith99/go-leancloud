package goleancloud

import (
	"bytes"
	"io"
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

func (cli *LeanClient) NewServiceContext() *ServiceContext {
	return &ServiceContext{
		LeanClient: cli,
	}
}

func (cli *LeanClient) NewRequest(method, apiUrl string, buf []byte) (*http.Request, error) {
	var body io.Reader
	if buf != nil {
		body = bytes.NewBuffer(buf)
	}

	req, err := http.NewRequest(method, apiUrl, body)
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
