package model

import "time"

// ArticleInfo 定义文章结构体
// 这个结构体中，没有定义文章内容字段，因为文章内容是大文本，而我们的文章列表功能，不需要使用文章内容信息，所以这里没有定义文章内容字段
// 这样可以效率高一点
type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Summary      string    `db:"summary"` // 文章摘要
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CreateTime   time.Time `db:"create_time"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
}

// ArticleDetail 定义文章详情页的实体，详情页需要使用文章内容，所以这个结构体将文章内容字段添加进去
type ArticleDetail struct {
	ArticleInfo
	Content  string `db:"content"`
	Category        // 相关文章需要使用当前文章的分类，所以我们将分类结构体也嵌套进来
}

// ArticleRecord 用于文章上下页
type ArticleRecord struct {
	ArticleInfo
	Category
}
