package api

import (
	"bxj_gorm_job_02/model"
	"bxj_gorm_job_02/model/response"

	"github.com/gin-gonic/gin"
)

type CommentApi struct {
}

func (api *CommentApi) Add(c *gin.Context) {
	var comment model.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.FailWithMessage("参数错误", c, err)
	}
	_, err = commentService.Add(&comment)
	if err != nil {
		response.FailWithMessage("添加失败", c, err)
	}
	response.OkWithData(comment, c)
}
func (api *CommentApi) GetCommentList(c *gin.Context) {
	posts, err := commentService.GetCommentList(c.Query("postId"))
	if err != nil {
		response.FailWithMessage("获取失败", c, err)
	}
	response.OkWithData(posts, c)
}
