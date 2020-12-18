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
	Token     string
	Stime     string
	Nonce     string
	Signature string
	Referer   string
}

// NewConnInfo 实例化连接信息
func NewConnInfo() *ConnInfo {
	return &ConnInfo{}
}

// ConnectionGdas 建立 Gdas 连接
func (c *ConnInfo) ConnectionGdas() (resp *http.Response, err error) {

	url := fmt.Sprintf("https://%v:%v%v", c.Addr, c.Port, c.Path)
	// 设置请求头
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("token", c.Token)
	req.Header.Set("stime", c.Stime)
	req.Header.Set("nonce", c.Nonce)
	req.Header.Set("signature", c.Signature)
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

// GetToken 获取 Token
func (c *ConnInfo) GetToken() (err error) {
	// 设置 json 格式的 request body
	jsonReqBody := []byte(`{"userName":"system","passWord":"d153850931040e5c81e1c7508ded25f5f0ae76cb57dc1997bc343b878946ba23"}`)
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
		return err
	}
	jsonRespBody, err := simplejson.NewJson(respBody)
	if err != nil {
		return err
	}
	// fmt.Printf("本次响应的 Body 为：%v\n响应中的 result 字段为：%v\n", string(body), js.Get("result"))
	c.Token, _ = jsonRespBody.Get("token").String()
	return
}
