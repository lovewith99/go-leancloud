package goleancloud

import (
	"bytes"
	"encoding/json"
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

func NewLeanClient(appId, appKey, masterKey string, options ...func(*LeanClient)) *LeanClient {
	cli := &LeanClient{
		AppId:     appId,
		AppKey:    appKey,
		MasterKey: masterKey,
		Endpoint:  "https://api.leancloud.cn",
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	for _, f := range options {
		f(cli)
	}

	return cli
}

func GetLeanClient(appId, appKey, masterKey string) *LeanClient {
	return NewLeanClient(appId, appKey, masterKey)
}

func (cli *LeanClient) SetEndpoint(endpoint string) {
	cli.Endpoint = endpoint
}

func (cli *LeanClient) NewServiceContext() *ServiceContext {
	return &ServiceContext{
		LeanClient: cli,
	}
}

func (cli *LeanClient) SetReqHeader(req *http.Request) {
	req.Header.Set("X-LC-Id", cli.AppId)
	req.Header.Set("X-LC-Key", cli.AppKey)
	req.Header.Set("Content-Type", "application/json")
}

func (cli *LeanClient) SetReqMasterHeader(req *http.Request) {
	req.Header.Set("X-LC-Id", cli.AppId)
	req.Header.Set("X-LC-Key", cli.MasterKey+",master")
	req.Header.Set("Content-Type", "application/json")
}

func (cli *LeanClient) DoRequest(req *http.Request, response interface{}) error {

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&response)
}

func (cli *LeanClient) Push(p *PushBody) error {
	req, err := http.NewRequest(
		"POST",
		cli.Endpoint+"/1.1/push",
		bytes.NewBuffer(p.Buffer()),
	)
	if err != nil {
		return err
	}
	cli.SetReqMasterHeader(req)

	var resp map[string]interface{}
	return cli.DoRequest(req, &resp)
}
