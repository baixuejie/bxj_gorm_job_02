package router

import (
	api2 "bxj_gorm_job_02/api"
	"bxj_gorm_job_02/model/response"
	"bxj_gorm_job_02/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 创建路由
	r := gin.Default()
	//配置需要登录放行的路由组
	authApi := r.Group("/api")
	{
		// 使用 & 创建结构体指针
		authApi.POST("/user/register", (&api2.UserApi{}).Register)
		authApi.POST("/user/login", (&api2.UserApi{}).Login)
	}
	// 配置受保护的路由组
	api := r.Group("/api").Use(AuthMiddleware())
	{
		// 文章模块
		api.POST("/post/add", (&api2.PostApi{}).Add)
		api.GET("/post/getPostList", (&api2.PostApi{}).GetPostList)
		api.GET("/post/getPostDetail", (&api2.PostApi{}).GetPostDetail)
		api.POST("/post/update", (&api2.PostApi{}).Update)
		api.DELETE("/post/delete", (&api2.PostApi{}).Delete)

		//评论模块
		api.POST("/comment/add", (&api2.CommentApi{}).Add)
		api.GET("/comment/getCommentList", (&api2.CommentApi{}).GetCommentList)
	}
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": response.UNAUTHORIZED,
				"msg":  "请求未携带token",
			})
			c.Abort()
			return
		}
		isLogin, err := utils.CheckToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": response.UNAUTHORIZED,
				"msg":  "无效的token",
			})
			c.Abort()
			return
		}
		if !isLogin {
			//项目中此处应跳转到登录页面
			c.JSON(http.StatusOK, gin.H{
				"code": response.UNAUTHORIZED,
				"msg":  "请先登录",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
