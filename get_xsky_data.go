package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	simplejson "github.com/bitly/go-simplejson"
)

// GetUsed 获取用户已使用空间
func GetUsed(c *ConnInfo) (u float64, err error) {
	// 查询 cluster 信息
	c.Path = "v1/cluster"
	resp, err := c.ConnectionXSky()
	defer resp.Body.Close()

	// 处理 Response Body,使用 simplejson 库获取响应 Body 的 *Json 类型数据
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	jsonRespBody, err := simplejson.NewJson(respBody)
	if err != nil {
		return
	}

	// 获取用户使用量
	used, _ := jsonRespBody.Get("cluster").Get("samples").GetIndex(0).Get("used_kbyte").Float64()
	fmt.Printf("当前用户已经使用了%vKiB\n", int64(used))
	u = used / 1024 / 1024 / 1024
	u, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", u), 64)
	return
}
