package mapper

import (
	"LiYZBlog/models"
	"LiYZBlog/models/dto"
	"LiYZBlog/models/vo"
	"LiYZBlog/util"
	"gorm.io/gorm"
)

func GetListByTitleAndCategoryId(title string, categoryId int) ([]models.Blog, error) {
	var blogs []models.Blog
	err := models.Db.Debug().Model(new(models.Blog)).
		Where("title like ? and category_id=?", "%"+title+"%", categoryId).
		Preload("Category").
		Find(&blogs).
		Error
	if err != nil {
		return nil, err
	}
	return blogs, err
}

func GetSearchBlogListByQueryAndIsPublished(query string) ([]vo.SearchBlog, error) {
	var blogs []vo.SearchBlog
	err := models.Db.Table("blog").
		Select("id,title,content").
		Where("is_published=? and password=? and content like ?", "true", "''", "%"+query+"%").
		Scan(&blogs).
		Error
	if err != nil {
		return nil, err
	}
	return blogs, err
}

func GetIdAndTitleList() ([]models.Blog, error) {
	var blogs []models.Blog
	err := models.Db.Model(new(models.Blog)).
		Select("id,title").
		Order("created_at desc").
		Error
	if err != nil {
		return nil, err
	}
	return blogs, err
}

func GetNewBlogListByIsPublished() ([]vo.NewBlog, error) {
	var blogs []vo.NewBlog
	err := models.Db.Model(new(models.Blog)).
		Select("id,title,password").
		Order("created_at desc").
		Scan(&blogs).
		Error
	if err != nil {
		return nil, err
	}
	return blogs, err
}

func GetBlogInfoListByIsPublished() ([]vo.BlogInfo, error) {
	var blogs []vo.BlogInfo
	err := models.Db.Table("blog").
		Select("blog.id,blog.title,blog.description,blog.is_top,blog.created_at as create_time,blog.views,blog.words,blog.read_time,blog.password,category.category_name").
		Joins("left join category on blog.category_id = category.id").
		Where("blog.is_published=true").
		Scan(&blogs).
		Error
	if err != nil {
		return nil, err
	}
	return blogs, err
}

func GetBlogInfoListByCategoryNameAndIsPublished(categoryName string) ([]vo.BlogInfo, error) {
	var blogs []vo.BlogInfo
	err := models.Db.Table("category c").
		Select("b.id,b.title,b.description,b.is_top,b.create_time,b.views,b.words,b.read_time,b.password,c.category_name").
		Joins("left join blog as b on b.category_id = c.id").
		Where("c.category_name = ? and and b.is_published = true", categoryName).
		Scan(&blogs).
		Error
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func GetBlogInfoListByTagNameAndIsPublished(tagName string) ([]models.Blog, error) {
	// tag表中通过tagName拿到tagId
	var tag models.Tag
	err := models.Db.Model(new(models.Tag)).
		Where("tag_name=?", tagName).
		First(&tag).
		Error
	if err != nil {
		return nil, err
	}
	tagId := tag.ID

	// blog_tag表中，通过tagId拿到blog_ids
	var blogTags []models.BlogTag
	err = models.Db.Model(new(models.BlogTag)).
		Where("tag_id=?", tagId).
		Find(&blogTags).
		Error
	if err != nil {
		return nil, err
	}
	var blogIds []int64
	for _, blogTag := range blogTags {
		blogIds = append(blogIds, blogTag.BlogId)
	}

	// 在blog表中通过blog_ids拿到blogs，利用外键加载Category
	var blogs []models.Blog
	err = models.Db.Model(new(models.Blog)).
		Where("is_published=? and id in ?", "true", blogIds).
		Preload("Category").
		Find(&blogs).
		Error
	if err != nil {
		return nil, err
	}
	return blogs, err
}

func GetGroupYearMonthByIsPublished() ([]string, error) {
	var date []string
	err := models.Db.Table("blog").
		Select("date_format(created_at, \"%Y年%m月\") as date").
		Where("is_published = ?", "true").
		Group("date_format(created_at, \"%Y年%m月\")").
		Order("date_format(created_at, \"%Y年%m月\") desc").
		Pluck("date", &date).
		Error
	if err != nil {
		return nil, err
	}
	return date, err
}

func GetArchiveBlogListByYearMonthAndIsPublished(yearMonth string) ([]vo.ArchiveBlog, error) {
	var archiveBlog []vo.ArchiveBlog
	err := models.Db.Table("blog").
		Select("id, title, password, date_format(create_time, \"%d日\") as day").
		Where("date_format(create_time, \"%Y年%m月\") = ? and is_published = true", yearMonth).
		Order("created_at desc").
		Scan(&archiveBlog).
		Error
	if err != nil {
		return nil, err
	}
	return archiveBlog, err
}

func GetRandomBlogListByLimitNumAndIsPublishedAndIsRecommend(limitNum int) ([]vo.RandomBlog, error) {
	var randomBlog []vo.RandomBlog
	err := models.Db.Table("blog").
		Select("id, title, password, create_time, first_picture").
		Where("is_published = true and is_recommend = true").
		Order("rand()").
		Limit(limitNum).
		Scan(&randomBlog).
		Error
	if err != nil {
		return nil, err
	}
	return randomBlog, err
}

func GetBlogViewsList() ([]dto.BlogView, error) {
	var blogView []dto.BlogView
	err := models.Db.Table("blog").
		Select("id, views").
		Scan(&blogView).
		Error
	if err != nil {
		return nil, err
	}
	return blogView, err
}

func DeleteBlogById(id int64) (int64, error) {
	var cnt int64
	err := models.Db.Where("id=?", id).Delete(models.Blog{}).Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func DeleteBlogTagByBlogId(blogId int64) (int64, error) {
	var cnt int64
	err := models.Db.Where("blog_id = ?", blogId).Delete(models.BlogTag{}).Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func SaveBlog(blog dto.Blog) (int64, error) {
	data := models.Blog{
		Title:            blog.Title,
		FirstPicture:     blog.FirstPicture,
		Content:          blog.Content,
		Description:      blog.Description,
		IsPublished:      util.IfUnit8(blog.Published, 1, 0),
		IsRecommend:      util.IfUnit8(blog.Recommend, 1, 0),
		IsAppreciation:   util.IfUnit8(blog.Appreciation, 1, 0),
		IsCommentEnabled: util.IfUnit8(blog.CommentEnabled, 1, 0),
		Views:            blog.Views,
		Words:            blog.Words,
		ReadTime:         blog.ReadTime,
		CategoryId:       int64(blog.Category.ID),
		IsTop:            util.IfUnit8(blog.Top, 1, 0),
		Password:         blog.Password,
		UserId:           int64(blog.User.ID),
	}

	var cnt int64
	err := models.Db.Create(&data).Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func SaveBlogTag(blogId int64, tagId int64) (int64, error) {
	blogTag := models.BlogTag{
		BlogId: blogId,
		TagId:  tagId,
	}

	var cnt int64
	err := models.Db.Create(&blogTag).Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func UpdateBlogRecommendById(blogId int64, recommend bool) (int64, error) {
	var cnt int64
	err := models.Db.Table("blog").
		Where("id=?", blogId).
		Update("is_recommend", util.IfUnit8(recommend, 1, 0)).
		Count(&cnt).
		Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func UpdateBlogVisibilityById(blogId int64, bv dto.BlogVisibility) (int64, error) {
	var blog = models.Blog{
		Model:            gorm.Model{ID: uint(blogId)},
		IsAppreciation:   util.IfUnit8(bv.Appreciation, 1, 0),
		IsRecommend:      util.IfUnit8(bv.Recommend, 1, 0),
		IsCommentEnabled: util.IfUnit8(bv.CommentEnabled, 1, 0),
		IsTop:            util.IfUnit8(bv.Top, 1, 0),
		IsPublished:      util.IfUnit8(bv.Published, 1, 0),
		Password:         bv.Password,
	}
	var cnt int64
	err := models.Db.Updates(&blog).Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func UpdateBlogTopById(blogId int64, top bool) (int64, error) {
	blog := models.Blog{
		Model: gorm.Model{
			ID: uint(blogId),
		},
		IsTop: util.IfUnit8(top, 1, 0),
	}
	var cnt int64
	err := models.Db.Updates(&blog).Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func UpdateViews(blogId int64, views int) (int64, error) {
	blog := models.Blog{
		Model: gorm.Model{
			ID: uint(blogId),
		},
		Views: views,
	}
	var cnt int64
	err := models.Db.Updates(&blog).Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func GetBlogById(id int64) (models.Blog, error) {
	var blog models.Blog
	err := models.Db.Table("blog b").
		Select("b.id,b.title,b.first_picture,b.content,b.description,b.is_recommend,b.is_published,b.is_appreciation,b.is_comment_enabled,b.is_top,b.create_time,b.update_time,b.views,b.words,b.read_time,b.password,c.id as category_id,c.category_name,bt.tag_id  as tag_id,t.tag_name as tag_name,t.color").
		Joins("left join category as c on b.category_id = c.id").
		Joins("left join blog_tag as bt on b.id = bt.blog_id").
		Joins("left join tag as t on bt.tag_id = t.id").
		Where("b.id=?", id).
		First(&blog).
		Error
	if err != nil {
		return models.Blog{}, err
	}
	return blog, err
}

func GetTitleByBlogId(id int64) (string, error) {
	var title string
	err := models.Db.Table("blog").
		Select("title").
		Pluck("title", &title).
		Error
	if err != nil {
		return "", err
	}
	return title, err
}

func GetBlogByIdAndIsPublished(id int64) (vo.BlogDetail, error) {
	var blogDetail vo.BlogDetail
	err := models.Db.Table("blog b").
		Select("b.id,b.title,b.content,b.is_appreciation,b.is_comment_enabled,b.is_top,b.create_time,b.update_time,b.views,b.words,b.read_time,b.password,c.category_name,t.tag_name as tag_name,t.color").
		Joins("left join category as c on b.category_id = c.id").
		Joins("left join blog_tag as bt on b.id = bt.blog_id").
		Joins("left join tag as t on bt.tag_id = t.id").
		Where("b.id=? and b.is_published=true", id).
		Scan(&blogDetail).
		Error
	if err != nil {
		return vo.BlogDetail{}, err
	}
	return blogDetail, err
}

func GetBlogPassword(blogId string) (string, error) {
	var password string
	err := models.Db.Table("blog").
		Select("password").
		Pluck("password", &password).
		Error
	if err != nil {
		return "", err
	}
	return password, err
}

func UpdateBlog(blog dto.Blog) (int64, error) {
	var cnt int64
	data := models.Blog{
		Model: gorm.Model{
			ID: uint(blog.Id),
		},
		Title:            blog.Title,
		FirstPicture:     blog.FirstPicture,
		Content:          blog.Content,
		Description:      blog.Description,
		IsPublished:      util.IfUnit8(blog.Published, 1, 0),
		IsRecommend:      util.IfUnit8(blog.Recommend, 1, 0),
		IsAppreciation:   util.IfUnit8(blog.Appreciation, 1, 0),
		IsCommentEnabled: util.IfUnit8(blog.CommentEnabled, 1, 0),
		Views:            blog.Views,
		Words:            blog.Words,
		ReadTime:         blog.ReadTime,
		CategoryId:       int64(blog.Category.ID),
		IsTop:            util.IfUnit8(blog.Top, 1, 0),
		Password:         blog.Password,
	}
	err := models.Db.Save(&data).Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func CountBlog() (int64, error) {
	var cnt int64
	err := models.Db.Table("blog").Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func CountBlogByIsPublished() (int64, error) {
	var cnt int64
	err := models.Db.Table("blog").Where("is_published = true").Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func CountBlogByCategoryId(categoryId int64) (int64, error) {
	var cnt int64
	err := models.Db.Table("blog").Where("category_id=?", categoryId).Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func CountBlogByTagId(tagId int64) (int64, error) {
	var cnt int64
	err := models.Db.Table("blog_tag").Where("tag_id = ?").Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func GetCommentEnabledByBlogId(blogId int64) (bool, error) {
	var isCommentEnabled uint8
	err := models.Db.Table("blog").
		Select("is_comment_enabled").
		Where("id=?", blogId).
		Pluck("is_comment_enabled", &isCommentEnabled).
		Error
	if err != nil {
		return false, err
	}
	return isCommentEnabled == 1, err
}

func GetPublishedByBlogId(blogId int64) (bool, error) {
	var isPublish uint8
	err := models.Db.Table("blog").
		Select("is_published").
		Where("id=?", blogId).
		Pluck("is_published", &isPublish).
		Error
	if err != nil {
		return false, err
	}
	return isPublish == 1, err
}

func GetCategoryBlogCountList() ([]vo.CategoryBlogCount, error) {
	var data []vo.CategoryBlogCount
	err := models.Db.Table("blog").
		Select("category_id as id, count(category_id) as blog_count").
		Group("category_id").
		Scan(&data).
		Error
	if err != nil {
		return nil, err
	}
	return data, err
}
