package goleancloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
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
	req, err := cli.NewRequest("GET", api, nil)
	if err != nil {
		return nil, err
	}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r GetServiceConvResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	if r.Error != "" {
		return &r, errors.New(r.Error)
	}

	return &r, nil
}

// PostService 创建服务号
func (cli *ServiceContext) PostService(name string) (*PostServiceConvResponse, error) {
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

// Pub 给所有订阅者发布消息
func (cli *ServiceContext) Pub(request *ServiceConvBroadcastRequest) (*ServiceConvBroadcastResponse, error) {
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
