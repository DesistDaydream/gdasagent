package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
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
	req.Header.Set("referer", c.Referer)

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
	jsonStr := []byte(`{"userName": "system","passWord": "d153850931040e5c81e1c7508ded25f5f0ae76cb57dc1997bc343b878946ba23"}`)
	fmt.Println("认证信息为：", bytes.NewBuffer(jsonStr))

	// 忽略证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// 建立连接
	url := fmt.Sprintf("https://%v:%v/v1/login", c.Addr, c.Port)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("referer", "https://172.38.30.192/gdas")

	// 获取响应
	resp, err := (&http.Client{Transport: tr}).Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	bodystr := string(body)

	fmt.Println(bodystr)
	// c.Token = bodystr
	return
}
