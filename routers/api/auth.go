package api

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	//获取url中username和password
	username := c.Query("username")
	password := c.Query("password")
	valid := validation.Validation{}
	//注意：定义valid变量，作用是存储上面结构体valid标签验证的结果。
	//验证通过，valid的error字段为空，否则包含所有验证失败的信息
	//定义结构体时就初始化  将获取到的信息传入结构体中（可能信息不符合规范）
	a := auth{Username: username, Password: password}
	//判断结构体是否合法，需要传入一个结构体的指针（容易忘记）
	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	//code初始化永远是无效的参数
	code := e.INVALID_PARAMS
	if ok {
		//传入用户名和密码判断数据库中是否存在
		isExist := models.CheckAuth(username, password)
		if isExist {
			//数据库中存在的话，产生token（一个字符串）（需要判断错误）
			token, err := util.GenerateToken(username, password)
			//判断错误
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}

		} else { //数据库中不存在
			code = e.ERROR_AUTH
		}

	} else {
		//如果结构体检验失败，则遍历err，然后打印err的键值对（err是map类型）
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    data,
	})

}
