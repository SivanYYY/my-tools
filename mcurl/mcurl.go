package mcurl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	GET  = "GET"
	POST = "POST"
)

func setHeader(req *http.Request, header *map[string]string) *http.Request {
	for k, v := range *header {
		req.Header.Set(k, v)
	}
	return req
}

// Get
// @description "http get请求"
// @param url "请求链接"
// @param header "头信息"
// @param clientTimeOut "请求超时时间"
// @return string "返回数据"
func Get(url string, header *map[string]string, clientTimeOut time.Duration) ([]byte, error) {
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		return nil, err
	}
	if header != nil {
		setHeader(req, header)
	}
	client := http.Client{
		Timeout: clientTimeOut,
	}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code is %d", rsp.StatusCode))
	}
	res, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PostJson
// @description "http get请求"
// @param url "请求链接"
// @param header "头信息"
// @param data "请求数据"
// @param clientTimeOut "请求超时时间"
// @return string "返回数据"
func PostJson(url string, header *map[string]string, data interface{}, clientTimeOut time.Duration) ([]byte, int, error) {
	jec, err := json.Marshal(data)
	if err != nil {
		return nil, 0, err
	}
	buffer := bytes.NewBuffer(jec)
	req, err := http.NewRequest(POST, url, buffer)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	if header != nil {
		for k, v := range *header {
			req.Header.Set(k, v)
		}
	}
	client := http.Client{
		Timeout: clientTimeOut,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, resp.StatusCode, errors.New(resp.Status)
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return result, resp.StatusCode, nil
}
