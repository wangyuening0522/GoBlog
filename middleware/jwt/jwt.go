package jwt

import (
	"blog/pkg/e"
	"blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// gin中间件，检验token有效性
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.Query("token")
		//判断是否存在token
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			//解码token
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				//判断是否超过过期时间
				//time.now().unix()  现在unix时间
				//claims.expiresat  设定的过期unix时间
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}

		}
		//果如code不是成功（其他一切的错误），均返回前端一段data数据
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			//停止进程
			c.Abort()
			return
		}
		//如果是成功的话，就继续
		c.Next()

	}

}
