package main

import (
	"fmt"
)

// 从 API 获取数据

// // Data 存储空间属性
// type Data struct {
// 	MaxAvailable string
// 	MaxSalable   string
// 	Used         string
// }

// // NewData 实例化 Data
// func (d *Data) NewData(ma string, ms string, u string) *Data {
// 	return &Data{
// 		MaxAvailable: ma,
// 		MaxSalable:   ms,
// 		Used:         u,
// 	}
// }

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
	// 待开发

	// 将数据写入到文件中
	// 待开发
}
