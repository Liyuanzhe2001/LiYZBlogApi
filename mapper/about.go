package mapper

import (
	"LiYZBlog/models"
)

func GetList() ([]models.About, error) {
	var abouts []models.About
	err := models.Db.Model(new(models.About)).Find(&abouts).Error
	return abouts, err
}

func UpdateAbout(nameEn string, value string) (int64, error) {
	var cnt int64
	err := models.Db.
		Model(new(models.About)).
		Where("name_en = ?", nameEn).
		Update("value", value).
		Count(&cnt).
		Error
	return cnt, err
}

func GetAboutCommentEnabled() (string, error) {
	var about *models.About
	err := models.Db.
		Select("value").
		Where("name_en = ?", "commentEnabled").
		First(&about).
		Error
	return about.Value, err
}
