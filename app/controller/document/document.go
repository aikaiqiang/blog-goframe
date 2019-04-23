package ctl_document

import (
	"gf-blog/app/library/document"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
)

func Index(r *ghttp.Request)  {
	if r.IsAjaxRequest(){

	}
}




// 处理ajax请求
func serveMarkdownAjax(r *ghttp.Request) {
	r.Response.WriteJson(g.Map{
		"code": 1,
		"msg":  "",
		"data": lib_document.GetMarkdown(r.Get("path", "index")),
	})
}