package api

import (
	"context"
	// "fmt"
    // "github.com/gogf/gf/v2/net/ghttp"
    "github.com/gogf/gf/v2/frame/g"

	core "github.com/rainxue/goimapi/core"
)

type ConversationController struct {

}

type ConvIdReq struct {
	g.Meta `path:"/v1/conversations/:Id" method:"get"`
	Id uint32
}

type ConvRes struct {
	*core.Conversation
	// Reply string `dc:"Reply content"`
	Reply string
}


func (c *ConversationController) Get(ctx context.Context, req *ConvIdReq) (res *ConvRes, err error) {// ghttp.Request   core.Conversation
	res = &ConvRes{
		Conversation: &core.Conversation{
			Id: req.Id,
			ConvType: "group",
			Title:"测试群1",
			// members []uint32
			// ext map
			CreateBy:2,
		},
		Reply:"xxx",
	}
	g.Log().Debugf(ctx, `receive say: %+v`, req)
	if(req.Id>100) {
		panic("www500 panic error!")
	}
	// res = &ConvRes{
	// 	Reply: fmt.Sprintf(`Hi %s`, req.Id),
	// }
	return 
	// return new(core.Conversation{})
}

// func (*Conversations) List(r *ghttp.Request) ([]core.Conversation, error) {
// 	return []
// }
// func (*Conversations) User(r *ghttp.Request) ([]*df_goods_sku.Entity, error) {

// }

