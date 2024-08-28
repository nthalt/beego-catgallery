package main

import (
	_ "beego-catgallery/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.SetStaticPath("/static", "static")
	beego.Run()
}

