package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Data 存储空间属性
type Data struct {
	MaxAvailable float64 `json:"MaxAvailable"`
	MaxSalable   float64 `json:"MaxSalable"`
	Used         float64 `json:"Used"`
}

// NewData 实例化 Data
func NewData() *Data {
	return &Data{
		MaxAvailable: 0,
		MaxSalable:   0,
		Used:         0,
	}
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
	if err := c.GetGdasToken(); err != nil {
		fmt.Printf("获取 Gdas Token 失败，失败原因：%v\n", err)
	}
	fmt.Printf("获取到的 Gdas Token 为：%v\n", c.Token)

	// 计算数据
	d := NewData()
	var err error
	if d.MaxAvailable, err = GetMaxAvailable(c); err != nil {
		fmt.Printf("获取最大可用存储空间数据失败，原因：%v", err)
	}
	d.MaxSalable = d.MaxAvailable
	if d.Used, _ = GetUsed(c); err != nil {
		fmt.Printf("获取用户已使用容量失败，原因：%v", err)
	}
	fmt.Printf("当前最大可用容量为 %vTiB\n", d.MaxAvailable)
	fmt.Printf("当前最大可售卖容量为 %vTiB\n", d.MaxSalable)
	fmt.Printf("当前已使用容量为 %vTiB\n", d.Used)

	// 将数据写入到文件中
	// 将结构体数据转为 JSON 数据
	result, _ := json.Marshal(d)
	fmt.Println("将要写入文件中的数据：", string(result))
	// 写入文件
	ioutil.WriteFile("/tmp/info.json", result, 0644)
}
