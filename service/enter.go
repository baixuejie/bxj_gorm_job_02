package service

type ServiceGroup struct {
	UserService
	CommentService
	PostService
}

var ServiceGroupApp = new(ServiceGroup)
