package main

import (
	"kala-v2/dao"
	"kala-v2/routers"
)

func main() {
	// 创建数据库
	// CREATE DATABASE kala;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	// 程序退出前关闭数据库连接
	defer dao.Close()
	// 注册路由
	r := routers.SetupRouter()
	// 启动服务
	r.Run(":80")
}
