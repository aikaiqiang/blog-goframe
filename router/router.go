package router

import (
    "gf-blog/app/controller/hello"
    "github.com/gogf/gf/g"
)

// 统一路由注册.
func init() {
    g.Server().BindHandler("/", ctl_hello.Handler)
}
