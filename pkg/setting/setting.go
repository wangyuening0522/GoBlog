package setting

//实现了将配置文件中的数据读取到全局变量的功能
import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSecret    string
)

// init函数会在程序启动时自动执行
func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}
func LoadBase() {
	//由于run-mode在全局区内，所以section是“空”
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

}
func LoadServer() {
	//获取的分区存到section中
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	//sec.Key("READ_TIMEOUT") 表示获取配置文件中 server 部分的 READ_TIMEOUT 键所对应的值
	//MustInt(60) 表示将其转换为整数类型，如果转换失败则返回默认值 60。
	//然后将这个整数值转换为 time.Duration 类型，最后乘以 time.Second 得到一个以秒为单位的时间段。  60是默认值
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// 关于一些疑问：READ_TIMEOUT = 60 WRITE_TIMEOUT = 60是什么
// 读取超时时间指的是服务器在接收到请求后，等待客户端发送请求的最长时间，超过，服务器关闭
// 写入超时时间指的是服务器向客户端发送响应的最长时间
// 总结：这些参数是在HTTP服务器启动时从配置文件中读取出来，并用于配置HTTP服务器的参数，从而让HTTP服务器能够正常地监听端口并处理请求。
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app':%v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)

}

//如果配置文件中存在 JWT_SECERT 键，则将其对应的值作为JWT的秘钥；否则，将使用 !@)*#)!@U#@*!@!) 作为默认秘钥。
//进行类型转换原因：从不同的数据源（如配置文件、数据库等）读取的数据类型可能不一致
