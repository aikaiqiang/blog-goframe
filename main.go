package main

import (
	_ "gf-blog/boot"
	_ "gf-blog/router"
	"github.com/gogf/gf/g"
)

func main() {
	g.Server().Run()
}
