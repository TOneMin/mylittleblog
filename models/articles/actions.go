package articles

var articleList []Article

//GetAllArticles get all the articles
func GetAllArticles() []Article {
	articleList = []Article{
		{ID: 1, Title: "Article 1", Body: "Article 1 body"},
		{ID: 2, Title: "Article 2", Body: "Article 2 body"},
		{ID: 3, Title: "Article 3", Body: "Article 3 body"},
		{ID: 4, Title: "Article 4", Body: "Article 4 body"},
	}

	return articleList
}

// GetArticle get specific article by id
func GetArticle(id int) *Article {
	for _, article := range articleList {
		if article.ID == id {
			return &article
		}
	}
	return nil
}
