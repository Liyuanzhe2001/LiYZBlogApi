package dto

import (
	"LiYZBlog/models"
	"time"
)

type Blog struct {
	Id             int64     `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`           // 文章标题
	FirstPicture   string    `json:"first_picture,omitempty"`   // 文章首图，用于随机文章展示
	Content        string    `json:"content,omitempty"`         // 文章正文
	Description    string    `json:"description,omitempty"`     //描述
	Published      bool      `json:"published,omitempty"`       //公开或私密
	Recommend      bool      `json:"recommend,omitempty"`       //推荐开关
	Appreciation   bool      `json:"appreciation,omitempty"`    //赞赏开关
	CommentEnabled bool      `json:"comment_enabled,omitempty"` //评论开关
	Top            bool      `json:"top,omitempty"`             //是否置顶
	CreateTime     time.Time `json:"create_time"`               //创建时间
	UpdateTime     time.Time `json:"update_time"`               //更新时间
	Views          int       `json:"views,omitempty"`           //浏览次数
	Words          int       `json:"words,omitempty"`           //文章字数
	ReadTime       int       `json:"read_time,omitempty"`       //阅读时长(分钟)
	Password       string    `json:"password,omitempty"`        //密码保护

	User     models.User     `json:"user"`               //文章作者(因为是个人博客，也可以不加作者字段，暂且加上)
	Category models.Category `json:"category"`           //文章分类
	Tags     []models.Tag    `json:"tags,omitempty"`     // 文章标签
	Cate     interface{}     `json:"cate,omitempty"`     // 页面展示层传输的分类对象：正常情况下为 字符串 或 分类id
	TagList  []interface{}   `json:"tag_list,omitempty"` //页面展示层传输的标签对象：正常情况下为 List<Integer>标签id 或 List<String>标签名
}

type BlogPassword struct {
	BlogId   int64  `json:"blog_id,omitempty"`
	Password string `json:"password,omitempty"`
}

type BlogView struct {
	Id   int64 `json:"id,omitempty"`
	View int   `json:"view,omitempty"`
}

type BlogVisibility struct {
	Appreciation   bool   `json:"appreciation,omitempty"`    //赞赏开关
	Recommend      bool   `json:"recommend,omitempty"`       //推荐开关
	CommentEnabled bool   `json:"comment_enabled,omitempty"` //评论开关
	Top            bool   `json:"top,omitempty"`             //是否置顶
	Published      bool   `json:"published,omitempty"`       //公开或私密
	Password       string `json:"password,omitempty"`        //密码保护
}

type Comment struct {
	Id              int64     `json:"id,omitempty"`
	Nickname        string    `json:"nickname,omitempty"`          //昵称
	Email           string    `json:"email,omitempty"`             //邮箱
	Content         string    `json:"content,omitempty"`           //评论内容
	Avatar          string    `json:"avatar,omitempty"`            //头像(图片路径)
	CreateTime      time.Time `json:"create_time"`                 //评论时间
	Website         string    `json:"website,omitempty"`           //个人网站
	Ip              string    `json:"ip,omitempty"`                //评论者ip地址
	Published       bool      `json:"published,omitempty"`         //公开或回收站
	AdminComment    bool      `json:"admin_comment,omitempty"`     //博主回复
	Page            int       `json:"page,omitempty"`              //0普通文章，1关于我页面
	Notice          bool      `json:"notice,omitempty"`            //接收邮件提醒
	ParentCommentId int64     `json:"parent_comment_id,omitempty"` //父评论id
	BlogId          int64     `json:"blog_id,omitempty"`           //所属的文章id
	Qq              string    `json:"qq,omitempty"`                //如果评论昵称为QQ号，则将昵称和头像置为QQ昵称和QQ头像，并将此字段置为QQ号备份
}

type Friend struct {
	Id          int64  `json:"id,omitempty"`
	Nickname    string `json:"nickname,omitempty"`    //昵称
	Description string `json:"description,omitempty"` //描述
	Website     string `json:"website,omitempty"`     //站点
	Avatar      string `json:"avatar,omitempty"`      //头像
	Published   bool   `json:"published,omitempty"`   //公开或隐藏
}

type LoginInfo struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Form struct {
	Id           string `json:"id,omitempty"`
	IsBot        bool   `json:"is_bot,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	Username     string `json:"username,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

type Chat struct {
	Id        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	Username  string `json:"username,omitempty"`
	Type      string `json:"type,omitempty"`
}

type Message struct {
	MessageId string `json:"message_id,omitempty"`
	Form      Form   `json:"form"`
	Chat      Chat   `json:"chat"`
	Date      string `json:"date,omitempty"`
	Text      string `json:"text,omitempty"`
}

type TgMessage struct {
	UpdateId string  `json:"update_id,omitempty"`
	Message  Message `json:"message"`
}

type UserAgentDTO struct {
	Os      string `json:"os,omitempty"`
	Browser string `json:"browser,omitempty"`
}

type VisitLogRemark struct {
	Content string `json:"content,omitempty"` // 访问内容
	Remark  string `json:"remark,omitempty"`  // 备注
}

type VisitLogUuidTime struct {
	Uuid string    `json:"uuid,omitempty"`
	Time time.Time `json:"time"`
	Pv   int       `json:"pv,omitempty"`
}
