package gdas

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/DesistDaydream/gdasagent/config"
	simplejson "github.com/bitly/go-simplejson"
)

// Gdas is
type Gdas struct {
	Addr   string
	Port   int
	Path   string
	Method string
	Token  string
	resp   *http.Response
}

// Connection is
func (g *Gdas) Connection() (err error) {
	url := fmt.Sprintf("https://%v:%v/%v", g.Addr, g.Port, g.Path)
	// 设置请求头
	req, _ := http.NewRequest(g.Method, url, nil)
	req.Header.Set("token", g.Token)
	req.Header.Set("referer", fmt.Sprintf("https://%v:%v/gdas", g.Addr, g.Port))

	// 忽略证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 建立连接
	if g.resp, err = (&http.Client{Transport: tr}).Do(req); err != nil {
		panic(err)
	}
	return
}

// GetToken is
func (g *Gdas) GetToken() (err error) {
	// 设置 json 格式的 request body
	jsonReqBody := []byte("{\"userName\":\"" + config.YamlConfig.Gdas.Username + "\",\"passWord\":\"" + config.YamlConfig.Gdas.Password + "\"}")
	// 设置 URL
	url := fmt.Sprintf("https://%v:%v/v1/login", g.Addr, g.Port)
	// 设置 Request 信息
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonReqBody))
	req.Header.Add("referer", fmt.Sprintf("https://%v:%v/gdas", g.Addr, g.Port))
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
	g.Token, _ = jsonRespBody.Get("token").String()
	return
}
