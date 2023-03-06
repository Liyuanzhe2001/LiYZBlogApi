package mapper

import (
	"LiYZBlog/models"
)

func GetListByTitleAndCategoryId(title string, categoryId int) ([]models.Blog, error) {
	var blogs []models.Blog
	err := models.Db.Model(new(models.Blog)).
		Where("title like ? and category_id=?", "%"+title+"%", categoryId).
		Preload("Category").
		Find(&blogs).
		Error
	if err != nil {
		return nil, err
	}
	return blogs, err
}
