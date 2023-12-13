package mcurl

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	GET  = "GET"
	POST = "POST"
)

// Get
// @description "http get请求"
// @param url "请求链接"
// @param header "头信息"
// @return string "返回数据"
func Get(url string, header *map[string]string) ([]byte, error) {
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		return nil, err
	}
	if header != nil {
		setHeader(req, header)
	}
	rsp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code is %d", rsp.StatusCode))
	}
	res, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func setHeader(req *http.Request, header *map[string]string) *http.Request {
	for k, v := range *header {
		req.Header.Set(k, v)
	}
	return req
}
