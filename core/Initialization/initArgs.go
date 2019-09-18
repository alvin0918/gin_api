package Initialization

import "flag"

var (
	ConfFile string // 配置文件路径
)

// 解析命令行参数
func InitArgs() {
	// master -config ./master.json -xxx 123 -yyy ddd
	// master -h
	flag.StringVar(&ConfFile, "c", "./app.json", "指定配置文件路径")
	flag.Parse()
}

