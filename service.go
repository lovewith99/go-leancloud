package goleancloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ServiceContext struct {
	*LeanClient
}

// GetService 查询服务号
func (cli *ServiceContext) GetService(name string, skip, limit int) (*GetServiceConvResponse, error) {
	p := url.Values{}
	p.Add("where", fmt.Sprintf(`{"name": "%s"}`, name))
	if skip > 0 {
		p.Add("skip", strconv.Itoa(skip))
	}
	if limit > 0 {
		p.Add("limit", strconv.Itoa(limit))
	}

	api := cli.Endpoint + "/1.2/rtm/service-conversations?" + p.Encode()
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return nil, err
	}
	cli.SetReqMasterHeader(req)

	var resp GetServiceConvResponse
	if err = cli.DoRequest(req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("%d:%s", resp.Code, resp.Error)
	}

	return &resp, nil
}

// PostService 创建服务号
func (cli *ServiceContext) PostService(name string) (*PostServiceConvResponse, error) {
	req, err := http.NewRequest(
		"POST",
		cli.Endpoint+"/1.2/rtm/service-conversations",
		strings.NewReader(fmt.Sprintf(`{"name": "%s"}`, name)),
	)
	if err != nil {
		return nil, err
	}

	var resp PostServiceConvResponse
	cli.SetReqMasterHeader(req)
	if err = cli.DoRequest(req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("%d:%s", resp.Code, resp.Error)
	}

	return &resp, nil
}

// Pub 给所有订阅者发布消息
func (cli *ServiceContext) Pub(request *ServiceConvBroadcastRequest) (*ServiceConvBroadcastResponse, error) {
	buf, _ := json.Marshal(request)
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/1.2/rtm/service-conversations/%s/broadcasts", cli.Endpoint, request.ConvId),
		bytes.NewBuffer(buf),
	)
	if err != nil {
		return nil, err
	}
	cli.SetReqMasterHeader(req)

	var resp ServiceConvBroadcastResponse
	if err = cli.DoRequest(req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("%d:%s", resp.Code, resp.Error)
	}

	return &resp, nil
}

// Subscribe 订阅
func (cli *ServiceContext) Subscribe(convId, clientId string) error {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/1.2/rtm/service-conversations/%s/subscribers", cli.Endpoint, convId),
		strings.NewReader(fmt.Sprintf(`{"client_id":"%s"}`, clientId)),
	)
	if err != nil {
		return nil
	}

	cli.SetReqMasterHeader(req)

	var resp map[string]interface{}
	return cli.DoRequest(req, &resp)
}

// Unsubscribe 取消订阅
func (cli *ServiceContext) Unsubscribe(convId, clientId string) error {
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/1.2/rtm/service-conversations/%s/subscribers/%s", cli.Endpoint, convId, clientId),
		nil,
	)
	if err != nil {
		return err
	}

	cli.SetReqMasterHeader(req)

	var resp map[string]interface{}
	return cli.DoRequest(req, &resp)
}
