package go_leancloud

import (
	"net/http"
	"strings"
	"io/ioutil"
)

type LeanClient struct {
	AppId     string
	AppKey    string
	MasterKey string
	*http.Client
}

func GetLeanClient(appId, appKey, masterKey string) *LeanClient {
	return &LeanClient{
		AppId:     appId,
		AppKey:    appKey,
		MasterKey: masterKey,
		Client:    http.DefaultClient,
	}
}

func (cli *LeanClient) NewRequest(method, apiUrl, buf string) (*http.Request, error) {
	req, err := http.NewRequest(method, apiUrl, strings.NewReader(buf))
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-LC-Id", cli.AppId)
	req.Header.Set("X-LC-Key", cli.AppKey)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (cli *LeanClient) Push(p *PushBody) error {
	req, err :=  cli.NewRequest("POST", "https://api.leancloud.cn/1.1/push", p.toString())
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
