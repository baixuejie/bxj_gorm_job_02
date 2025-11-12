package service

import (
	"bxj_gorm_job_02/config"
	"bxj_gorm_job_02/model"
)

type PostService struct {
}

func (p *PostService) GetPostList() ([]model.Post, error) {
	var posts []model.Post
	err := config.DB.Find(&posts).Error
	return posts, err
}
func (p *PostService) GetPostDetail(id uint) (model.Post, error) {
	var post model.Post
	err := config.DB.Where("id = ?", id).First(&post).Error
	return post, err
}
func (p *PostService) AddPost(post *model.Post) (*model.Post, error) {
	err := config.DB.Create(post).Error
	return post, err
}
func (p *PostService) DeletePost(post *model.Post) (*model.Post, error) {
	err := config.DB.Delete(post).Error
	return post, err
}
func (p *PostService) UpdatePost(post *model.Post) (*model.Post, error) {
	err := config.DB.Save(post).Error
	return post, err
}
