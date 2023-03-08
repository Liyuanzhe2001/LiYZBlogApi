package vo

import (
	"LiYZBlog/models"
	"time"
)

type ArchiveBlog struct {
	id       int64
	title    string
	day      string
	password string
	privacy  bool
}

type Badge struct {
	title   string
	url     string
	subject string
	value   string
	color   string
}

type BlogDetail struct {
	id             int64
	title          string
	content        string
	appreciation   bool      //赞赏开关
	commentEnabled bool      //评论开关
	top            bool      //是否置顶
	createTime     time.Time //创建时间
	updateTime     time.Time // 更新时间
	views          int       // 浏览次数
	words          int       // 文章字数
	readTime       int       //阅读时长(分钟)
	password       string    //密码保护

	category models.Category //文章分类
	tags     []models.Tag    //文章标签
}

type BlogIdAndTitle struct {
	id    int64
	title string
}

type BlogInfo struct {
	id          int64
	title       string    // 文章标题
	description string    //描述
	createTime  time.Time //创建时间
	views       int       //浏览次数
	words       int       //文章字数
	readTime    int       //阅读时长(分钟)
	top         bool      //是否置顶
	password    string    //文章密码
	privacy     bool      //是否私密文章

	category models.Category //文章分类
	tags     []models.Tag    // 文章标签
}

type CategoryBlogCount struct {
	id    int64
	name  string //分类名
	value int    //分类下博客数量
}

type Copyright struct {
	title    string
	siteName string
}

type Favorite struct {
	title   string
	content string
}

type Friend struct {
	nickname    string //昵称
	description string //描述
	website     string //站点
	avatar      string //头像
}

type Introduction struct {
	avatar   string
	name     string
	github   string
	telegram string
	qq       string
	bilibili string
	netease  string
	email    string

	rollText  []string
	favorites []Favorite
}

type NewBlog struct {
	id       int64
	title    string
	password string
	privacy  bool
}

type PageComment struct {
	id                    int64
	nickname              string    //昵称
	content               string    //评论内容
	avatar                string    //头像(图片路径)
	createTime            time.Time //评论时间
	website               string    //个人网站
	adminComment          bool      //博主回复
	parentCommentId       string    //父评论id
	parentCommentNickname string    //父评论昵称

	replyComments []PageComment //回复该评论的评论
}

type PageResult struct {
	totalPage int           //总页数
	list      []interface{} // 数据
}

type RandomBlog struct {
	id           int64
	title        string    // 文章标题
	firstPicture string    //文章首图，用于随机文章展示
	createTime   time.Time //创建时间
	password     string    //文章密码
	privacy      bool      //是否私密文章
}

type SearchBlog struct {
	id      int64
	title   string
	content string
}

type TagBlogCount struct {
	id    int64
	name  string //标签名
	value int    //标签下博客数量
}
