package e

var MsgFlags = map[int]string{
	//错误码                         错误信息     相互映射
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "token生成失败",
	ERROR_AUTH:                     "Token错误",
}

func GetMsg(code int) string {
	//如果 MsgFlags 映射表中存在错误码为 code 的错误信息，则将其赋值给 msg 变量，ok为ture
	//。如果 MsgFlags 映射表中不存在错误码为 code 的错误信息，则 msg 变量的值为空字符串，ok 变量为 false。
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	//返回引射表中error对应的错误信息
	return MsgFlags[ERROR]
}

//同一级文件夹下，可以不引用而直接使用另一个go文件下的全局变量，声明的常量
