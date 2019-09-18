package Initialization

import "runtime"

// 初始化线程数量
func InitEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
