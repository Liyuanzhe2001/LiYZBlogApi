package vo

import (
	"LiYZBlog/models"
	"time"
)

type ArchiveBlog struct {
	Id       int64  `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Day      string `json:"day,omitempty"`
	Password string `json:"password,omitempty"`
	Privacy  bool   `json:"privacy,omitempty"`
}

type Badge struct {
	Title   string `json:"title,omitempty"`
	Url     string `json:"url,omitempty"`
	Subject string `json:"subject,omitempty"`
	Value   string `json:"value,omitempty"`
	Color   string `json:"color,omitempty"`
}

type BlogDetail struct {
	Id             int64     `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	Appreciation   bool      `json:"appreciation,omitempty"`    //赞赏开关
	CommentEnabled bool      `json:"comment_enabled,omitempty"` //评论开关
	Top            bool      `json:"top,omitempty"`             //是否置顶
	CreateTime     time.Time `json:"create_time"`               //创建时间
	UpdateTime     time.Time `json:"update_time"`               // 更新时间
	Views          int       `json:"views,omitempty"`           // 浏览次数
	Words          int       `json:"words,omitempty"`           // 文章字数
	ReadTime       int       `json:"read_time,omitempty"`       //阅读时长(分钟)
	Password       string    `json:"password,omitempty"`        //密码保护

	Category models.Category `json:"category"`       //文章分类
	Tags     []models.Tag    `json:"tags,omitempty"` //文章标签
}

type BlogIdAndTitle struct {
	Id    int64
	Title string
}

type BlogInfo struct {
	Id          int64     `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`       // 文章标题
	Description string    `json:"description,omitempty"` //描述
	CreateTime  time.Time `json:"create_time"`           //创建时间
	Views       int       `json:"views,omitempty"`       //浏览次数
	Words       int       `json:"words,omitempty"`       //文章字数
	ReadTime    int       `json:"read_time,omitempty"`   //阅读时长(分钟)
	Top         bool      `json:"top,omitempty"`         //是否置顶
	Password    string    `json:"password,omitempty"`    //文章密码
	Privacy     bool      `json:"privacy,omitempty"`     //是否私密文章

	Category models.Category `json:"category"`       //文章分类
	Tags     []models.Tag    `json:"tags,omitempty"` // 文章标签
}

type CategoryBlogCount struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`  //分类名
	Value int    `json:"value,omitempty"` //分类下博客数量
}

type Copyright struct {
	Title    string `json:"title,omitempty"`
	SiteName string `json:"site_name,omitempty"`
}

type Favorite struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type Friend struct {
	Nickname    string `json:"nickname,omitempty"`    //昵称
	Description string `json:"description,omitempty"` //描述
	Website     string `json:"website,omitempty"`     //站点
	Avatar      string `json:"avatar,omitempty"`      //头像
}

type Introduction struct {
	Avatar   string `json:"avatar,omitempty"`
	Name     string `json:"name,omitempty"`
	Github   string `json:"github,omitempty"`
	Telegram string `json:"telegram,omitempty"`
	Qq       string `json:"qq,omitempty"`
	Bilibili string `json:"bilibili,omitempty"`
	Netease  string `json:"netease,omitempty"`
	Email    string `json:"email,omitempty"`

	RollText  []string   `json:"roll_text,omitempty"`
	Favorites []Favorite `json:"favorites,omitempty"`
}

type NewBlog struct {
	Id       int64  `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Password string `json:"password,omitempty"`
	Privacy  bool   `json:"privacy,omitempty"`
}

type PageComment struct {
	Id                    int64     `json:"id,omitempty"`
	Nickname              string    `json:"nickname,omitempty"`                //昵称
	Content               string    `json:"content,omitempty"`                 //评论内容
	Avatar                string    `json:"avatar,omitempty"`                  //头像(图片路径)
	CreateTime            time.Time `json:"create_time"`                       //评论时间
	Website               string    `json:"website,omitempty"`                 //个人网站
	AdminComment          bool      `json:"admin_comment,omitempty"`           //博主回复
	ParentCommentId       string    `json:"parent_comment_id,omitempty"`       //父评论id
	ParentCommentNickname string    `json:"parent_comment_nickname,omitempty"` //父评论昵称

	ReplyComments []PageComment `json:"reply_comments,omitempty"` //回复该评论的评论
}

type PageResult struct {
	TotalPage int           //总页数
	List      []interface{} // 数据
}

type RandomBlog struct {
	Id           int64     `json:"id,omitempty"`
	Title        string    `json:"title,omitempty"`         // 文章标题
	FirstPicture string    `json:"first_picture,omitempty"` //文章首图，用于随机文章展示
	CreateTime   time.Time `json:"create_time"`             //创建时间
	Password     string    `json:"password,omitempty"`      //文章密码
	Privacy      bool      `json:"privacy,omitempty"`       //是否私密文章
}

type SearchBlog struct {
	Id      int64  `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type TagBlogCount struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`  //标签名
	Value int    `json:"value,omitempty"` //标签下博客数量
}
