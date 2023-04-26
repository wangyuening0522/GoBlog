package models

import (
	"blog/pkg/setting"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

// gorm.Open()函数用于打开一个数据库连接，返回一个*gorm.DB实例和一个错误示例，使用db来接受返回的*gorm.DB实例
var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	//拿到配置文件database区
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	//先拿到配置文件中的值

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	//fmt.Sprintf()用于格式化字符串，将参数按照制定的格式转换成一个数据库连接字符串
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	//这里用到的变量都在前面被声明为string，并且都被赋值（在配置文件中）
	if err != nil {
		log.Println(err)
	}
	//处理默认表名，匿名函数DefaultTableNameHandler 接收两个参数（db *gorm.DB和默认表名）
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName //返回的新表名是自定义的表前缀和默认表明拼接起来的（通常是子谷川常量或者从配置文件中读取的值）
	}
	//true  gorm会将结构体名称复数到单数（Users->user）
	db.SingularTable(true)
	//ture输出sql日志
	db.LogMode(true)
	//设置连接池空闲连接最大数量 在 GORM 的默认设置中，空闲连接的数量为 2。
	db.DB().SetMaxIdleConns(10)
	//设置连接池中打开连接的最大数量。在 GORM 的默认设置中，打开连接的数量为 0，即没有限制。
	db.DB().SetMaxOpenConns(100)
}

// 总结：通过限制连接池中打开连接的数量，可以避免因连接过多导致数据库崩溃或者性能下降的问题。
// 开启 SQL 日志功能可以方便地跟踪和调试应用程序中的数据库操作。
func CloseDB() {
	defer db.Close()
}

//连接池相关：
/*
1.要素：
最小连接数：连接池中维护的最小连接数，即初始化时创建的连接数。
最大连接数：连接池中维护的最大连接数，即中能够同时存在的最大连接数。
空闲连接数：连接池中当前未被使用的连接数。
活动连接数：连接池中当前正在被使用的连接数。
2.概念：
连接池维护着一组可复用的数据库连接，应用程序可以从连接池中获取连接，并在使用完成后将连接归还给连接池。
3.好处：
复用已存在的连接，减少连接数据库的开销
*/
