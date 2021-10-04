package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"test_crud/models"
)

// ArticleHandler  represent the httphandler for article
type ArticleHandler struct {
	AUsecase models.ArticleUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticleHandler(e *echo.Echo, articleUsecase models.ArticleUsecase) {
	handler := &ArticleHandler{
		AUsecase: articleUsecase,
	}
	e.GET("/articles", handler.GetArticles)
	//e.POST("/articles", handler.Store)
	//e.GET("/articles/:id", handler.GetByID)
	//e.DELETE("/articles/:id", handler.Delete)
}

// FetchArticle will fetch the article based on given params
func (a *ArticleHandler) GetArticles(c echo.Context) error {
	listArr := a.AUsecase.GET()
	return c.JSON(http.StatusOK, listArr)
}