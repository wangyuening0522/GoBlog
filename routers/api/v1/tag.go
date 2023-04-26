package v1

import (
	"blog/models"
	"blog/pkg/e"
	"blog/pkg/setting"
	"blog/pkg/util"
	validation "github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// 获取多个文章标签
func GetTags(c *gin.Context) {
	//主要实现从前端传来的url地址中拿到参数，并且将其作为筛选条件在数据库中查询和返回结果
	name := c.Query("name")
	//url中查询name，查到赋值给name，未查到返回“”字符串
	maps := make(map[string]interface{})
	//用maps存储查询条件，这里使name和state
	data := make(map[string]interface{})
	//用data存储符合条件的列表
	if name != "" {
		//maps数据的添加
		maps["name"] = name
	}
	//默认state为-1，表示不对状态进行过滤
	var state int = -1
	//url中查询state，查到赋值给arg，未查到返回“”字符串
	if arg := c.Query("state"); arg != "" {
		//满足存在state，往下走
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS
	//models文件夹下面的taga.go文件里的gettags函数和gettagstotal函数，引用进来包就可以调用
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	//拿到数据列表                    数据提取初始位置，数据查询个数，查询条件
	data["total"] = models.GetTagTotal(maps)
	//拿到数据总数     查询条件，返回符合条件数据个数
	c.JSON(http.StatusOK, gin.H{
		"code": code,           //状态码
		"msg":  e.GetMsg(code), //拿到状态码对应信息
		"data": data,           //拿到数据data  map
	})
}

// 新增文章标签
func AddTag(c *gin.Context) {
	//拿到url中name，state，createdby参数
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")
	//写入规范验证
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态至于允许为0或1")
	//初始话状态码是invalid参数  400
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	//这里如果走的是else则不该百年code，所以可能会报错，写上下面的c.json就正确了
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		//错误码对应报错信息
		"msg": e.GetMsg(code),
		//返回空对象
		"data": make(map[string]string),
	})

}

// 修改文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")
	valid := validation.Validation{}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Required(id, "id").Message("ID不嫩为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	//初始化code 400
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		data := make(map[string]interface{})
		data["modified_by"] = modifiedBy
		if name != "" {
			data["name"] = name
		}
		if state != -1 {
			data["state"] = state
		}
		models.EditTag(id, data)
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// 删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
