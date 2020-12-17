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
	fmt.Println("认证信息为：", bytes.NewBuffer(jsonReqBody))
	// 设置 URL
	url := fmt.Sprintf("https://%v:%v/v1/login", c.Addr, c.Port)
	fmt.Printf("URL 为：%v\n", url)
	// 设置 Request 信息
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonReqBody))
	// req.Header.Add("referer", fmt.Sprintf("https://%v:%v/gdas", c.Addr, c.Port))
	// req.Header.Add("Cookie", "JSESSIONID=A3A81CF6835CA074957F2B1E838CB5A9")
	// req.Header.Add("Postman-Token", "e5759bb9-3c5c-40b4-911a-cf971dc61b95")
	// req.Header.Add("Content-Length", "101")
	// req.Header.Add("Host", "172.38.30.192:8003")

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

	fmt.Printf("Response 信息为：%v\n：", resp)
	fmt.Printf("Request 信息为：%v\n", resp.Request)

	// 处理 Response Body,并获取 Token
	body, err := ioutil.ReadAll(resp.Body)
	js, err := simplejson.NewJson(body)
	fmt.Printf("本次响应的 Body 为：%v\n响应中的 result 字段为：%v\n", string(body), js.Get("result"))

	// c.Token = js.Get("result")

	return
}
