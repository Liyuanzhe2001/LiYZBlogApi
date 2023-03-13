package mapper

import (
	"LiYZBlog/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetCityVisitorList() ([]models.CityVisitor, error) {
	var data []models.CityVisitor
	err := models.Db.Table("city_visitor").
		Select("city, uv").
		Order("uv dessc").
		Find(&data).
		Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func SaveCityVisitor(cityVisitor models.CityVisitor) (int64, error) {
	var cnt int64
	err := models.Db.
		Create(&cityVisitor).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "city"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"uv": gorm.Expr("uv+?", cityVisitor.Uv)}),
		}).
		Count(&cnt).
		Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}
