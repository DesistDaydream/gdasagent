package main

import (
	"github.com/thedevsaddam/gojsonq/v2"
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
	c.GetToken()
	// 关闭连接

	// 计算数据
	// 待开发

	// 将数据写入到文件中
	// 待开发
	const json = `{"name":{"first":"Tom","last":"Hanks"},"age":61}`
	name := gojsonq.New().FromString(json).Find("name.first")
	println(name.(string)) // Tom
}
