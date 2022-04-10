package common
// 修改 GHTTP/ghttp_middleware_handler_response.go
import (
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/errors/gerror"
	// "github.com/gogf/gf/internal/intlog"
)

type DefaultHandlerResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// MiddlewareHandlerResponse is the default middleware handling handler response object and its error.
func MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()
	var (
		err         error
		res         interface{}
		internalErr error
	)
	res, err = r.GetHandlerResponse()
	if err != nil {
		code := gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		internalErr = r.Response.WriteJson(DefaultHandlerResponse{
			Code:    code.Code(),
			Message: err.Error(),
			Data:    nil,
		})
		if internalErr != nil {
			// intlog.Error(r.Context(), internalErr)
			// TODO: 
		}
		return
	}
	// internalErr = r.Response.WriteJson(DefaultHandlerResponse{
	// 	Code:    gcode.CodeOK.Code(),
	// 	Message: "",
	// 	Data:    res,
	// })
	internalErr = r.Response.WriteJson(res)
	if internalErr != nil {
		// intlog.Error(r.Context(), internalErr)
		// TODO: 
	}
}