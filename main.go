package main

import (
	_ "master-slave-go-demo/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

