package routers

import (
	"littleblog/models/articles"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitializeRouters initialize routers
func InitializeRouters(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello there")
	})
	router.GET("/articles", articles.GetAllArticlesHandler)
	router.GET("/article/:id", articles.GetArticleHandler)
}
