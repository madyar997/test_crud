package http

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"test_crud/models"
)

type ArticleHandler struct {
	AUsecase models.ArticleUsecase
}

func NewArticleHandler(e *echo.Echo, articleUsecase models.ArticleUsecase) {
	handler := &ArticleHandler{
		AUsecase: articleUsecase,
	}
	e.GET("/articles", handler.GetAll)
	e.POST("/articles", handler.Create)
	//e.GET("/articles/:id", handler.GetByID)
	e.DELETE("/articles/:id", handler.Delete)
	e.PUT("/articles", handler.Update)
}

func (a *ArticleHandler) GetAll(c echo.Context) error {
	listArr, err := a.AUsecase.Get()
	if err != nil {
		log.Fatal("error while getting article: ", err)
	}
	return c.JSON(http.StatusOK, listArr)
}

func (a *ArticleHandler) Create(c echo.Context) error {
	article := new(models.Article)
	if err := c.Bind(&article); err != nil {
		return err
	}
	err := a.AUsecase.Create(article)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Article has been successfully added")
}

func (a *ArticleHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Please specify the id parameter")
	}
	id := int64(idP)
	err = a.AUsecase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (a *ArticleHandler) Update(c echo.Context) error {
	article := new(models.Article)
	if err := c.Bind(&article); err != nil {
		return err
	}
	err := a.AUsecase.Update(article)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "Article has been successfully updated")
}
