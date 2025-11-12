package router

import (
	api2 "bxj_gorm_job_02/api"
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
		authApi.POST("/user/register", api2.UserApi{}.Register)
		authApi.POST("/user/login", api2.UserApi{}.Login)

	}
	// 配置受保护的路由组
	api := r.Group("/api")
	{
		// 用户模块
		api.POST("/user/getUserInfo", api2.UserApi{}.GetUserInfo)
		// 文章模块
		api.POST("/post/add", api2.PostApi{}.Add)
		api.POST("/post/getPostList", api2.PostApi{}.GetPostList)
		api.POST("/post/getPostDetail", api2.PostApi{}.GetPostDetail)
		api.POST("/post/update", api2.PostApi{}.Update)
		api.DELETE("/post/delete", api2.PostApi{}.Delete)

		//评论模块
		api.POST("/comment/add", api2.CommentApi{}.Add)
		api.POST("/comment/getCommentList", api2.CommentApi{}.GetCommentList)
		api.DELETE("/comment/delete", api2.CommentApi{}.Delete)
		api.POST("/comment/getCommentDetail", api2.CommentApi{}.GetCommentDetail)
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
				"code": -1,
				"msg":  "请求未携带token",
			})
			c.Abort()
			return
		}
		isLogin, err := utils.CheckToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "无效的token",
			})
		}
		if !isLogin {
			//项目中此处应跳转到登录页面
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "请先登录",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
