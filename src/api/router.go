package api

import (
    "github.com/gogf/gf/v2/net/ghttp"
    "github.com/gogf/gf/v2/frame/g"

    // api "github.com/rainxue/goimapi/api"
    // common "github.com/rainxue/goimapi/common"
)

func InitRouters() {
	s := g.Server()
    s.Use(ghttp.MiddlewareHandlerResponse)
    // s.Use(common.MiddlewareHandlerResponse)

    // conversation
    c := &ConversationController{}
    s.BindHandler("GET:/v1/conversations", func(r *ghttp.Request) {
		r.Response.Write("[]")
	})
    s.BindHandler("POST:/v1/conversations", func(r *ghttp.Request) {
		r.Response.Write("post.")
	})
    s.BindHandler("GET:/v1/conversations/:Id", c.Get)
	// conversation_group := s.Group("/v1/conversations")
	// // group.ALL("/all", func(r *ghttp.Request) {
	// // 	r.Response.Write("all")
	// // })
    // conversation_group.GET("/", func(r *ghttp.Request) {
	// 	r.Response.Write("[]")
	// })    
    // conversation_group.POST("/", func(r *ghttp.Request) {
	// 	r.Response.Write("post")
	// })
	// conversation_group.GET("/:id", func(r *ghttp.Request) {
	// 	r.Response.Write("get")
	// })


    // conversation member


    // conversation message


    // user message


    // s.BindHandler("/:name", func(r *ghttp.Request){
    //    r.Response.Writeln(r.Router.Uri)
    // })
    
	// s.SetPort(8998)
    s.Run()
}