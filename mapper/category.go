package mapper

import "LiYZBlog/models"

func GetCategoryList() ([]models.Category, error) {
	var categories []models.Category
	err := models.Db.Table("category").
		Select("id,category_name").
		Find(&categories).
		Error
	if err != nil {
		return nil, err
	}
	return categories, err
}

func GetCategoryNameList() ([]models.Category, error) {
	var categories []models.Category
	err := models.Db.Table("category").
		Select("category_name").
		Order("id desc").
		Find(&categories).
		Error
	if err != nil {
		return nil, err
	}
	return categories, err
}

func SaveCategory(category models.Category) (int64, error) {
	var cnt int64
	err := models.Db.Table("category").
		Update("category_name", category.CategoryName).
		Count(&cnt).
		Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func GetCategoryById(id int64) (models.Category, error) {
	var category models.Category
	err := models.Db.Table("category").
		Select("id,category_name").
		Where("id=?", id).
		First(&category).
		Error
	if err != nil {
		return models.Category{}, err
	}
	return category, err
}

func GetCategoryByName(name string) (models.Category, error) {
	var category models.Category
	err := models.Db.Table("category").
		Select("id, category_name").
		Where("category_name=?", name).
		First(&category).
		Error
	if err != nil {
		return models.Category{}, err
	}
	return category, err
}

func DeleteCategoryById(id int64) (int64, error) {
	var cnt int64
	err := models.Db.
		Where("id=?", id).
		Delete(models.Category{}).
		Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func UpdateCategory(category models.Category) (int64, error) {
	var cnt int64
	err := models.Db.Table("category").
		Where("id=?", category.ID).
		Update("category_name=?", category.CategoryName).
		Count(&cnt).
		Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}
