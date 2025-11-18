package api

import (
	"bxj_gorm_job_02/model"
	"bxj_gorm_job_02/model/response"

	"github.com/gin-gonic/gin"
)

type PostApi struct {
}

func (api *PostApi) Add(c *gin.Context) {
	var post model.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		response.FailWithMessage("参数错误", c, err)
	}
	_, err = postService.AddPost(&post)
	if err != nil {
		response.FailWithMessage("添加失败", c, err)
	}
	response.OkWithMessage("添加成功", c)
}

func (api *PostApi) GetPostDetail(c *gin.Context) {
	post, err := postService.GetPostDetail(c.Query("id"))
	if err != nil {
		response.FailWithMessage("获取失败", c, err)
	}
	response.OkWithData(post, c)
}

func (api *PostApi) GetPostList(c *gin.Context) {
	posts, err := postService.GetPostList()
	if err != nil {
		response.FailWithMessage("获取失败", c, err)
	}
	response.OkWithData(posts, c)
}

func (api *PostApi) Delete(c *gin.Context) {
	result, err := postService.DeletePost(c.Query("id"))
	if err != nil {
		response.FailWithMessage("删除失败", c, err)
	}
	response.OkWithData(result, c)

}

func (api *PostApi) Update(c *gin.Context) {
	var post model.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		response.FailWithMessage("参数错误", c, err)
	}
	result, err := postService.UpdatePost(&post)
	if err != nil {
		response.FailWithMessage("更新失败", c, err)
	}
	response.OkWithData(result, c)
}
