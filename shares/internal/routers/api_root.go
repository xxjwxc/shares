package routers

import (
	"net/http"

	"shares/internal/api"
	"shares/internal/service/analy"
	_ "shares/internal/service/analy"
	"shares/internal/service/shares"
	"shares/internal/service/weixin"
	proto "shares/rpc/shares"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/gmsec/micro/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/xxjwxc/ginrpc"
	"github.com/xxjwxc/public/dev"
	"github.com/xxjwxc/public/tools"
)

// OnInitRoot 初始化
func OnInitRoot(s server.Server, router gin.IRoutes, objs ...interface{}) {
	var args []interface{}
	w := new(weixin.Weixin)
	h := new(shares.Shares)
	a := new(analy.Analy)
	args = append(args, w, h, a)
	if s != nil {
		proto.RegisterWeixinServer(s, w) // 服务注册
		proto.RegisterSharesServer(s, h)
	}
	args = append(args, objs...)
	OnInitRouter(router, args...)
}

// OnInitRouter 默认初始化
func OnInitRouter(router gin.IRoutes, objs ...interface{}) {
	InitFunc(router)
	InitObj(router, objs...)
	init3rdGrpcHost()
}

// 初始化第三方host
func init3rdGrpcHost() {
	// micro.SetClientServiceName(oauth2.GetOauth2Name(), "gmsec.srv.oauth2")
}

// InitFunc 默认初始化函数
func InitFunc(router gin.IRoutes) {
	router.StaticFS("/file", http.Dir(tools.GetCurrentDirectory()+"/file")) //加载静态资源，一般是上传的资源，例如用户上传的图片
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	}) // 健康检查
	router.GET("/metrics", ginprom.PromHandler(promhttp.Handler())) // 添加grafana监控
}

// InitObj 初始化对象
func InitObj(router gin.IRoutes, objs ...interface{}) {
	base := ginrpc.New(ginrpc.WithCtx(api.NewAPIFunc), ginrpc.WithOutDoc(dev.IsDev()), ginrpc.WithDebug(dev.IsDev()),
		ginrpc.WithOutPath("internal/routers"), ginrpc.WithImportFile("rpc/common", "../apidoc/rpc/common"),
		ginrpc.WithBeforeAfter(&ginrpc.DefaultGinBeforeAfter{}))
	router.POST("/analy.deal_msg", base.HandlerFunc(analy.DealMsg))
	base.RegisterHandlerFunc(router, []string{"post", "get"}, "/analy.wx_token", analy.WxTokenSignature)

	base.OutDoc(true)
	base.Register(router, objs...) // 对象注册
}

// Cors 跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
