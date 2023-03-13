package mapper

import (
	"LiYZBlog/models"
	"LiYZBlog/models/dto"
	"LiYZBlog/models/vo"
	"LiYZBlog/util"
)

func GetListByPageAndParentCommentId(page int, blogId int64, parentCommentId int64) ([]models.Comment, error) {
	var data []models.Comment
	err := models.Db.Table("comment c").
		Select("c.id, c.nickname, c.email, c.content, c.avatar, c.create_time, c.ip, c.is_published, c.is_admin_comment, c.page,c.is_notice, c.parent_comment_id, c.website, c.blog_id, c.qq, b.title").
		Joins("left join blog as b on c.blog_id=b.id").
		Where("c.page=? and c.blog_id=? and c.parent_comment_id=?", page, blogId, parentCommentId).
		First(&data).
		Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func GetListByParentCommentId(parentCommentId int64) ([]models.Comment, error) {
	var data []models.Comment
	err := models.Db.Table("comment c").
		Select("c.id, c.nickname, c.email, c.content, c.avatar, c.create_time, c.ip, c.is_published, c.is_admin_comment, c.page,c.is_notice, c.parent_comment_id, c.website, c.qq, c.blog_id, b.title").
		Joins("left join blog as b on c.blog_id=b.id").
		Where("c.parent_comment_id=?", parentCommentId).
		First(&data).
		Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func GetPageCommentListByPageAndParentCommentId(page int, blogId int64, parentCommentId int64) ([]vo.PageComment, error) {
	var data []vo.PageComment
	err := models.Db.Table("comment c1").
		Select("c1.id, c1.nickname, c1.content, c1.avatar, c1.create_time, c1.is_admin_comment, c1.website,c1.parent_comment_id as parent_comment_id, c2.nickname as parent_comment_nickname").
		Joins("left join comment as c2 on c1.parent_comment_id=c2.id").
		Where("c1.page=? and c1.blog_id=? and c1.parent_comment_id=? and c1.is_published=true", page, blogId, parentCommentId).
		Order("c1.created_at desc").
		Scan(&data).
		Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func GetCommentById(id int64) (models.Comment, error) {
	var comment models.Comment
	err := models.Db.Table("comment c").
		Select("c.id, c.nickname, c.email, c.content, c.avatar, c.create_time, c.ip, c.is_published, c.is_admin_comment, c.page,c.is_notice, c.parent_comment_id, c.website, c.qq, c.blog_id, b.title").
		Joins("left join blog as b on c.blog_id=b.id").
		Where("c.id=?", id).
		First(&comment).
		Error
	if err != nil {
		return models.Comment{}, err
	}
	return comment, err
}

func UpdateCommentPublishedById(commentId int64, published bool) (int64, error) {
	var cnt int64
	err := models.Db.Table("comment").
		Where("id=?", commentId).
		Update("is_published", util.IfInterface(published, 1, 0)).
		Count(&cnt).
		Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func UpdateCommentNoticeById(commentId int64, notice bool) (int64, error) {
	var cnt int64
	err := models.Db.Table("comment").
		Where("id=?", commentId).
		Update("is_published", util.IfInterface(notice, 1, 0)).
		Count(&cnt).
		Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func DeleteCommentById(commentId int64) (int64, error) {
	var cnt int64
	err := models.Db.Where("id=?", commentId).Delete(models.Comment{}).Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func DeleteCommentsByBlogId(blogId int64) (int64, error) {
	var cnt int64
	err := models.Db.Where("blog_id=?", blogId).Delete(models.Comment{}).Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func UpdateComment(comment models.Comment) (int64, error) {
	var cnt int64
	err := models.Db.Table("comment").Save(&comment).Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func CountByPageAndIsPublished(page int, blogId int64, isPublished bool) (int64, error) {
	var cnt int64
	err := models.Db.Table("comment").
		Select("count(*) as cnt").
		Where("page=? and is_published=? and blog_id=?", page, isPublished, blogId).
		Pluck("cnt", &cnt).
		Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func CountComment() (int64, error) {
	var cnt int64
	err := models.Db.Table("comment").
		Where("count(*) as cnt").
		Pluck("cnt", &cnt).
		Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}

func SaveComment(comment dto.Comment) (int64, error) {
	var cnt int64
	err := models.Db.Save(&comment).Count(&cnt).Error
	if err != nil {
		return -1, err
	}
	return cnt, err
}
