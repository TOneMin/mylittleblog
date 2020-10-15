package articles

import (
	"littleblog/middlewares"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllArticlesHandler(c *gin.Context) {
	middlewares.CORSMiddleware()(c)
	articleList := GetAllArticles()
	c.JSON(http.StatusOK, articleList)
}

func GetArticleHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
	}
	article := GetArticle(id)
	c.JSON(http.StatusOK, article)
}
