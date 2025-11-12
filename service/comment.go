package service

import (
	"bxj_gorm_job_02/config"
	"bxj_gorm_job_02/model"
)

type CommentService struct {
}

func (s *CommentService) Add(c *model.Comment) (*model.Comment, error) {
	config.DB.Create(c)
	return c, nil
}
func (s *CommentService) GetCommentList(postId uint) ([]model.Comment, error) {
	config.DB.Where("post_id = ?", postId).Find(&model.Comment{})
	return nil, nil
}
func (s *CommentService) Delete(c *model.Comment) (*model.Comment, error) {
	config.DB.Delete(c)
	return nil, nil
}
func (s *CommentService) GetCommentDetail(id uint) (model.Comment, error) {
	var comment model.Comment
	config.DB.Where("id = ?", id).Find(&comment)
	return comment, nil
}
func (s *CommentService) Update(c *model.Comment) (*model.Comment, error) {
	config.DB.Save(c)
	return nil, nil
}
