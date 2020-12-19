package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
)

// ConnInfo 连接信息
type ConnInfo struct {
	Method string
	Addr   string
	Port   int
	Path   string
	// 认证信息
	Token   string
	Referer string
}

// NewConnInfo 实例化连接信息
func NewConnInfo() *ConnInfo {
	return &ConnInfo{
		Method:  "GET",
		Addr:    "",
		Port:    0,
		Path:    "",
		Token:   "",
		Referer: "",
	}
}

// ConnectionGdas 根据给定的 path 建立 Gdas 连接
func (c *ConnInfo) ConnectionGdas() (resp *http.Response, err error) {
	url := fmt.Sprintf("https://%v:%v/%v", c.Addr, c.Port, c.Path)
	// 设置请求头
	req, _ := http.NewRequest(c.Method, url, nil)
	req.Header.Set("token", c.Token)
	req.Header.Set("referer", fmt.Sprintf("https://%v:%v/gdas", c.Addr, c.Port))

	// 忽略证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 建立连接
	if resp, err = (&http.Client{Transport: tr}).Do(req); err != nil {
		panic(err)
	}
	return
}

// ConnectionXSky 根据给定的 path 建立 XSky 连接
func (c *ConnInfo) ConnectionXSky() (resp *http.Response, err error) {
	url := fmt.Sprintf("http://%v:%v/%v", c.Addr, c.Port, c.Path)
	// 设置请求头
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", "XMS_AUTH_TOKEN="+c.Token)

	// 忽略证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 建立连接
	if resp, err = (&http.Client{Transport: tr}).Do(req); err != nil {
		panic(err)
	}
	return
}

// GetGdasToken 获取 Token
func (c *ConnInfo) GetGdasToken() (err error) {
	// 设置 json 格式的 request body
	// json :=
	jsonReqBody := []byte("{\"userName\":\"" + yamlConfig.Gdas.Username + "\",\"passWord\":\"" + yamlConfig.Gdas.Password + "\"}")
	// 设置 URL
	url := fmt.Sprintf("https://%v:%v/v1/login", c.Addr, c.Port)
	// 设置 Request 信息
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonReqBody))
	req.Header.Add("referer", fmt.Sprintf("https://%v:%v/gdas", c.Addr, c.Port))
	req.Header.Add("Content-Type", "application/json")

	// 忽略证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// 发送 Request 并获取 Response
	resp, err := (&http.Client{Transport: tr}).Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 处理 Response Body,并获取 Token
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	jsonRespBody, err := simplejson.NewJson(respBody)
	if err != nil {
		return
	}
	// fmt.Printf("本次响应的 Body 为：%v\n响应中的 result 字段为：%v\n", string(body), js.Get("result"))
	c.Token, _ = jsonRespBody.Get("token").String()
	return
}

// GetXSkyToken 获取 XSky Token
func (c *ConnInfo) GetXSkyToken() (err error) {
	// 设置 json 格式的 request body
	// json :=
	jsonReqBody := []byte("{\"auth\":{\"name\":\"" + yamlConfig.Xsky.Username + "\",\"password\":\"" + yamlConfig.Xsky.Password + "\"}}")
	// jsonReqBody := []byte(`{"auth":{"name":"admin","password":"admin"}}`)
	// 设置 URL
	url := fmt.Sprintf("http://%v:%v/api/v1/auth/tokens:login", c.Addr, c.Port)
	// 设置 Request 信息
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonReqBody))
	req.Header.Add("Content-Type", "application/json")

	// 发送 Request 并获取 Response
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 处理 Response Body,并获取 Token
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	jsonRespBody, err := simplejson.NewJson(respBody)
	if err != nil {
		return
	}
	// fmt.Printf("本次响应的 Body 为：%v\n", string(respBody))
	c.Token, _ = jsonRespBody.Get("token").Get("uuid").String()
	return
}
