package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/DesistDaydream/gdasagent/config"
	simplejson "github.com/bitly/go-simplejson"
)

// GetMaxAvailable 获取最大可用的存储空间
func GetMaxAvailable(c *ConnInfo) (m float64, err error) {
	// 查询盘匣列表
	c.Path = config.YamlConfig.Gdas.Path
	resp, err := c.ConnectionGdas()
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

	// 使用 simplejson 库处理 *Json 类型数据,len() 统计数组中元素个数，即盘匣总数
	// 判断 rfidSts 值为 0 的盘匣就是未分配，循环所有盘匣的 rfidSts 的值，获取未分配的总数
	var undistributedCount float64
	for i := 0; i < len(jsonRespBody.Get("rfid").MustArray()); i++ {
		rfidSts, _ := jsonRespBody.Get("rfid").GetIndex(i).Get("rfidSts").Int64()
		// fmt.Println(test)
		// 如果 rfidSts 为 0，则计数器加1
		if rfidSts == 0 {
			undistributedCount = undistributedCount + 1
		}
	}
	fmt.Printf("当前共有 %v 个盘匣，其中 %v 个未分配\n", len(jsonRespBody.Get("rfid").MustArray()), undistributedCount)
	m = undistributedCount * 2.59
	m, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", m), 64)

	return
}
