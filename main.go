package main

import (
	_ "cool-admin-goframe/boot"
	_ "cool-admin-goframe/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
