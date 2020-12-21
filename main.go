package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/DesistDaydream/gdasagent/config"
)

// Data 要展示的数据应该具有的属性
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

// GetGdasData 从 Gdas 中获取数据
func GetGdasData(c *ConnInfo, d *Data) (err error) {
	// 设置基本连接参数
	c.Addr = config.YamlConfig.Gdas.IP
	c.Port = config.YamlConfig.Gdas.Port
	// 获取 Token
	if err = c.GetGdasToken(); err != nil {
		fmt.Printf("获取 Gdas Token 失败，失败原因：%v\n", err)
	}
	fmt.Printf("获取到的 Gdas Token 为：%v\n", c.Token)

	// 计算数据
	if d.MaxAvailable, err = GetMaxAvailable(c); err != nil {
		fmt.Printf("获取最大可用存储空间数据失败，原因：%v", err)
	}
	d.MaxSalable = d.MaxAvailable
	return
}

// GetXSkyData 从 XSky 中获取数据
func GetXSkyData(c *ConnInfo, d *Data) (err error) {
	// 设置基本连接参数
	c.Addr = config.YamlConfig.Xsky.IP
	c.Port = config.YamlConfig.Xsky.Port
	// 获取 Token
	if err = c.GetXSkyToken(); err != nil {
		fmt.Printf("获取 XSky Token 失败，失败原因：%v\n", err)
	}
	fmt.Printf("获取到的 XSky Token 为：%v\n", c.Token)

	// 计算数据
	if d.Used, err = GetUsed(c); err != nil {
		fmt.Printf("获取用户已使用容量失败，原因：%v\n", err)
	}
	return
}

// 赶工作品，有待优化
// 优化点1：获取 Token 操作最好放在 middleware 目录中，作为 中间件 来使用
// 优化点2：代码过于流水账，无 go 哲学思想的体现
// 优化点3：各个结构体属性需要精简
// 优化点4：感觉 connection.go 有问题，但是想不到在哪
func main() {
	c := NewConnInfo()
	d := NewData()
	// 通过文件读取参数
	// 在 yaml.go 中通过 config.yaml 文件获取数据

	// GetGdasData 从 Gdas 中获取数据
	GetGdasData(c, d)
	GetXSkyData(c, d)

	fmt.Printf("当前最大可用容量为 %vTiB\n", d.MaxAvailable)
	fmt.Printf("当前最大可售卖容量为 %vTiB\n", d.MaxSalable)
	fmt.Printf("当前已使用容量为 %vTiB\n", d.Used)

	// 将数据写入到文件中
	// 将结构体数据转为 JSON 数据
	result, _ := json.Marshal(d)
	// fmt.Println("将要写入文件中的数据：", string(result))
	// 写入文件
	ioutil.WriteFile("./info.json", result, 0644)
}
