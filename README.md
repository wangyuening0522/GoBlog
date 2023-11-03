# GoBlog
使用goland，gin，gorm完成的博客
# 编写的接口
1.列表
- 获取标签列表：GET("/tags”)
- 新建标签：POST("/tags”)
- 更新指定标签：PUT("/tags/:id”)
- 删除指定标签：DELETE("/tags/:id”)
2.文章
- 获取文章列表：GET("/articles”)
- 获取指定文章：POST("/articles/:id”)
- 新建文章：POST("/articles”)
- 更新指定文章：PUT("/articles/:id”)
- 删除指定文章：DELETE("/articles/:id”)
# 身份校验
使用JWT(JSON Web Token)进行来访者身份校验，防止接口信息被随意调用
