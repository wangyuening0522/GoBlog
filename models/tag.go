package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	//在数据库中查找maps类型数据，offset(偏移/起始位置)，limit(限制10个数据)，find（执行查询）将查询结果存储在tags中，tags被声明为数组结构提并且被返回
	//例如查询年龄在18到30岁之间的男性用户，并且从头开始返回前10条数据
	/*db.Where("age BETWEEN ? AND ?", 18, 30).
	Where("gender = ?", "male").
	Order("age ASC").年龄升序排序
	Offset(pageNum).
	Limit(pageSize).
	Find(&users)*/

	return
}
func GetTagTotal(maps interface{}) (count int) {
	//Model方法指定要操作的模型，count方法统计符合条件的记录数
	//db.Model(&Tag{})
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag) //可以用find，但是find返回所有符合数据
	//意思是：              db中哪里由name=name的标签，只拿目标标签的id属性，只拿第一个符合条件的标签，并且存储到提前声明的tag变量中
	/*例如：符合条件的标签返回的结果是
	tag.ID = 123
	tag.Name = ""  // name 字段被忽略，值为默认值 ""
	tag.State = 0  // state 字段被忽略，值为默认值 0*/
	//select("id")查询结果只包含id字段，其他字段忽略
	//where查询条件，name字段等于指定名称的标签
	//first(&tag)只查询一条记录，将结果存出在传入的tag变量中
	if tag.ID > 0 {
		return true
	}
	return false

}

// 添加三个属性
func AddTag(name string, state int, createdBy string) bool {
	//使用gorm orm库的creat方法像数据库中添加一个新的标签
	//db：数据库连接对象，gorm的db对象
	//create():创建并且添加，接收一个指向（要插入的数据结构）的指针
	//这里使用 & 取地址符来获取该对象的指针，以便传递给 Create 方法。
	db.Create(
		&Tag{
			Name:      name,
			State:     state,
			CreatedBy: createdBy,
		})
	return true
}

// 方法接收一个指向gorm.Scope结构体的指针作为参数    结构体包含当前记录的所有操作和信息
// 模型定义了一个Tag的结构体，并定义对应的数据库表的结构和字段
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	//setcolum设置  指定列（字段）为当前值  设置created_on为time.Now().Unix()
	//钩子方法，nil表示操作成功
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

/*
这属于gorm的Callbacks，可以将回调方法定义为模型结构的指针，
在创建、更新、查询、删除时将被调用，如果任何回调返回错误，gorm 将停止未来操作并回滚所有更改。
*/
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id=?", id).Updates(data)
	//使用Model方法获取数据库操作对象，关联一个Tag模型   查询条件，update方法传入一个map类型的data参数
	//update用于更新符合条件的记录的字段值
	/*
	   map中的某个字段值为nil，则相应的字段不会被更新
	   更新成功，则返回一个包含影响的记录数的整数值和一个error对象；
	   更新失败，则只返回一个error对象，表示更新过程中出现的错误。
	*/
	return true
}
func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id=?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}
