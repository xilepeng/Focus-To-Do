package main

import (
	"fmt"
	conf "to-do-list/config"
	"to-do-list/repository/db/dao"
	"to-do-list/routes"
)

func main() { // http://localhost:3000/swagger/index.html
	loading()
	// 转载路由 swag init -g common.go
	r := routes.NewRouter()
	fmt.Println("启动成功...")
	_ = r.Run(conf.HttpPort)
}

func loading() {
	// 从配置文件读入配置
	conf.Init()
	dao.MySQLInit()
}
