package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Data 存储空间属性
type Data struct {
	MaxAvailable float64
	MaxSalable   string
	Used         string
}

// NewData 实例化 Data
func NewData() *Data {
	return &Data{}
}

// 从 API 获取数据
func main() {
	c := NewConnInfo()
	// 获取一些请求参数
	// 待开发
	// 设置基本连接参数
	c.Addr = "172.38.30.192"
	c.Port = 8003

	// 获取 Token
	if err := c.GetToken(); err != nil {
		fmt.Printf("获取token失败，失败原因：%v\n", err)
	}
	fmt.Printf("获取到的 Token 为：%v\n", c.Token)

	// 计算数据
	d := NewData()
	d.MaxAvailable, _ = GetMaxAvailable(c)
	fmt.Printf("当前最大可用容量为 %vTiB\n", d.MaxAvailable)
	fmt.Printf("当前最大可售卖容量为 %vTiB\n", d.MaxAvailable)
	// 当前已使用容量待开发
	fmt.Printf("当前已使用容量为 %vTiB\n", d.Used)

	// 将数据写入到文件中
	// 将结构体数据转为 JSON 数据
	result, _ := json.Marshal(d)
	fmt.Println(string(result))
	// 写入文件
	ioutil.WriteFile("info.json", result, 0644)
}
