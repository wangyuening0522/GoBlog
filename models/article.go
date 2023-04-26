package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model             //继承model结构体。model中定义了id，createdat，updateat
	TagID      int    `json:"tag_id" gorm:"index"` //tagid用于存储文章对应的标签的id，并通过gorm：“index”指定为索引
	Tag        Tag    `json:"tag"`                 //能够达到Article、Tag关联查询的功能,嵌套struct  tag表示文章对应标签
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// 为了在创造和更新以前，手动设置createon和modifiedon的时间，而且createon和modifiedon四嵌套gorm model结构体获得的
// gorm的钩子函数，返回nil，表示函数没有发生任何错误
func (articel *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix()) //指定createdon这列为当前时间戳
	return nil
}
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

// 通过id来看是否有此文章
func ExitArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id=?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

// 拿到文章总数，使用gorm中模型查询和计数功能
func GetArticleTotal(maps interface{}) (count int) {
	//获取article对应的数据库表  使用where(maps)方法指定查询条件
	db.Model(&Article{}).Where(maps).Count(&count) //count执行查询，结果存储在count变量中
	return
}

// 获取符合条件的文章列表article    页码，每页显示的记录数，指定查询条件
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preloads("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return //offset从pagenum位置开始查询 limit只需要查询10个      使用Find(&articles)方法执行查询，并将结果存储在articles变量中。
}

// 传入id，拿到文章返回article结构体
func GetArticle(id int) (article Article) {
	db.Where("id=?", id).First(&article)     //返回第一个article数据，存入articel中
	db.Model(&article).Related(&article.Tag) //将符合条件的article和db关联，related查找于article关联的tag的所有记录
	return
}

//*****重要：
/*能够达到关联，首先是gorm本身做了大量的约定俗成
1.Article有一个结构体成员是TagID，就是外键。gorm会通过类名+ID 的方式去找到这两个类之间的关联关系
2.Article有一个结构体成员是Tag，就是我们嵌套在Article里的Tag结构体，我们可以通过Related进行关联查询
// 传入id和修改内容，编辑文章*/
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id=?", id).Updates(data) //model指定查询的模型类型
	return true
}

// 添加文章，传入一个map  命名为data
func AddArticle(data map[string]interface{}) bool {
	//向数据库中添加一条记录
	db.Create(&Article{
		//类型断言x.()返回bool值,将传入的data数据类型转换称article需要的类型和键
		//将data[key]转换为对应相符的类型
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true //全部成功，则返回真
}

// 删除文章  传入id，返回是否成功
func DeleteArticle(id int) bool {
	//删除一条articel记录，delete要求传一个结构体类型的零值
	db.Where("id=?", id).Delete(Article{})
	return true
}
