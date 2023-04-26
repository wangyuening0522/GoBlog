package routers

import (
	"blog/middleware/jwt"
	"blog/pkg/setting"
	"blog/routers/api"
	v1 "blog/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//*gin.Engine 是指向Engine 结构体类型的指针。
	//Engine 结构体用于管理 HTTP 请求的路由和中间件等。
	//gin.New类似gin.Default, 创建了一个新的 Engine 对象
	r := gin.New()
	r.Use(gin.Logger())
	//logger中间件，请求时打印在终端，提示程序员请求时间，路径，方式
	r.Use(gin.Recovery())
	//recovery中间件，捕捉panic错误并将错误信息输出到控制台，返回500错误给客户端
	gin.SetMode(setting.RunMode)
	//新增获取token的路由
	r.GET("/auth", api.GetAuth)
	// r.Group("/api/v1") 创建了一个路由分组，用于处理 API 版本 1 的请求。
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id")
	}
	//函数返回了创建的 Engine 对象 r,外部程序可以使用该对象来启动 Web 服务器和监听 HTTP 请求。
	return r
}
