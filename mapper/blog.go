package mapper

import (
	"LiYZBlog/models"
	"LiYZBlog/models/vo"
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

//func GetArchiveBlogListByYearMonthAndIsPublished(yearMonth string) ([]models.Blog, error) {
//	var blogs []models.Blog
//	models.Db.Model(new(models.Blog)).
//		Select("id, title, password")
//}
