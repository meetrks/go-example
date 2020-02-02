package v1

import (
	"gin-example/models"
	"gin-example/pkg/app"
	"gin-example/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestData(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

type PostBlogForm struct {
	Title       string `form:"title" valid:"Required;MaxSize(100)"`
	Description string `form:"description" valid:"Required;MaxSize(500)"`
}

func PostBlog(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form PostBlogForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	blog := models.Blog{
		Title:       form.Title,
		Description: form.Description,
	}
	blog.Add()
}

func BlogList(c *gin.Context) {
	appG := app.Gin{C: c}

	blogs := models.Blog{}
	res := blogs.GetAll()

	data := make(map[string]interface{})
	data["result"] = res
	data["count"] = 1

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
