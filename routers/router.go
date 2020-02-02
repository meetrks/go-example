package routers

import (
	"gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	router.GET("/test", v1.TestData)
	router.POST("/post_blog", v1.PostBlog)
	router.GET("/blog", v1.BlogList)

	router.Run()
}
