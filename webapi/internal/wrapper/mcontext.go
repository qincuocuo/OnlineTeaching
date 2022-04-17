package wrapper

import (
	"sync"
	"webapi/common"
	"webapi/middleware/jwts"

	"github.com/kataras/iris/v12"
)

type Context struct {
	iris.Context
	UserToken   *common.UserToken
	AllowAccess []int
	AllowWrite  []int
}

var contextPool = sync.Pool{New: func() interface{} {
	return &Context{}
}}

// 获取用户Token相关信息
func acquire(original iris.Context) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Context = original // set the context to the original one in order to have access to iris's implementation.
	ctx.UserToken = jwts.GetUserToken(original)
	if ctx.Context.Values().Get("AllowAccess") != nil {
		ctx.AllowAccess = ctx.Context.Values().Get("AllowAccess").([]int) // get user allow access relateid set
	} // reset the jwt token
	if ctx.Context.Values().Get("AllowWrite") != nil {
		ctx.AllowWrite = ctx.Context.Values().Get("AllowWrite").([]int) // get user can or want update res relateid set
	}
	return ctx
}

//用户Token信息传递
func release(ctx *Context) {
	ctx.UserToken = nil
	ctx.AllowAccess = ctx.AllowAccess[0:0]
	contextPool.Put(ctx)
}

func Handler(handler func(*Context)) iris.Handler {
	return func(original iris.Context) {
		ctx := acquire(original)
		handler(ctx)
		release(ctx)
	}
}
