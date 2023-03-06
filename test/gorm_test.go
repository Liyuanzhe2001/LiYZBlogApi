package test

import (
	"LiYZBlog/mapper"
	"LiYZBlog/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGormTest(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/liyzblog?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	//db.AutoMigrate(&models.About{})
	db.AutoMigrate(&models.Blog{})
	//db.AutoMigrate(&models.BlogTag{})
	//db.AutoMigrate(&models.Category{})
	//db.AutoMigrate(&models.CityVisitor{})
	//db.AutoMigrate(&models.Comment{})
	//db.AutoMigrate(&models.ExceptionLog{})
	//db.AutoMigrate(&models.Friend{})
	//db.AutoMigrate(&models.LoginLog{})
	//db.AutoMigrate(&models.Moment{})
	//db.AutoMigrate(&models.OperationLog{})
	//db.AutoMigrate(&models.ScheduleJob{})
	//db.AutoMigrate(&models.ScheduleJobLog{})
	//db.AutoMigrate(&models.SiteSetting{})
	//db.AutoMigrate(&models.Tag{})
	//db.AutoMigrate(&models.User{})
	//db.AutoMigrate(&models.VisitRecord{})
	//db.AutoMigrate(&models.Visitor{})
	//db.AutoMigrate(&models.VisitLog{})

	var data []*models.Blog
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range data {
		fmt.Printf("User  => %v \n", v)
	}
}

func TestAbout(t *testing.T) {
	abouts, err := mapper.GetList()
	if err != nil {
		fmt.Println("错误", err)
		return
	}
	for _, about := range abouts {
		fmt.Println(about.Value)
	}
	//about, err := mapper.UpdateAbout("1", "10")
	//enabled, err := mapper.GetAboutCommentEnabled()
	//if err != nil {
	//	fmt.Println("错误", err)
	//	return
	//}
	//fmt.Println(about)
	//fmt.Println(enabled)

}

func TestBlog(t *testing.T) {
	id, err := mapper.GetListByTitleAndCategoryId("标", 1)
	if err != nil {
		fmt.Println("错误")
	}
	fmt.Println(id[0])
	fmt.Println(id[0].Category.CategoryName)
}
