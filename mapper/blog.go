package mapper

import (
	"LiYZBlog/models"
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

func GetSearchBlogListByQueryAndIsPublished(query string) ([]models.Blog, error) {
	var blogs []models.Blog
	err := models.Db.Model(new(models.Blog)).
		Select("id,title,content").
		Where("is_published=? and password=? and content like ?", "true", "''", "%"+query+"%").
		Find(&blogs).
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

func GetNewBlogListByIsPublished() ([]models.Blog, error) {
	var blogs []models.Blog
	err := models.Db.Model(new(models.Blog)).
		Select("id,title,password").
		Order("created_at desc").
		Error
	if err != nil {
		return nil, err
	}
	return blogs, err
}

func GetBlogInfoListByIsPublished() ([]models.Blog, error) {
	var blogs []models.Blog
	err := models.Db.Model(new(models.Blog)).
		Preload("Category").
		Where("is_published=?", "true").
		Order("created_at desc").
		Error
	if err != nil {
		return nil, err
	}
	return blogs, err
}

func GetBlogInfoListByCategoryNameAndIsPublished(categoryName string) ([]models.Blog, error) {
	var categories []models.Category
	err := models.Db.Model(new(models.Category)).
		Select("id").
		Where("category_name=?", categoryName).
		Find(&categories).
		Error
	if err != nil {
		return nil, err
	}

	var ids []uint
	for _, category := range categories {
		ids = append(ids, category.ID)
	}

	var blogs []models.Blog
	err = models.Db.Model(new(models.Blog)).
		Preload("Category").
		Where("category_id in (?) and is_published=?", ids, "true").
		Find(&blogs).
		Error
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

//func GetBlogInfoListByTagNameAndIsPublished(tagName string) ([]models.Blog, error) {
//}
