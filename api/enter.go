package api

import "bxj_gorm_job_02/service"

var (
	userService    = service.ServiceGroupApp.UserService
	postService    = service.ServiceGroupApp.PostService
	commentService = service.ServiceGroupApp.CommentService
)
