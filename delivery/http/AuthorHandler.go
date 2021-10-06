package http

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"test_crud/models"
)

type AuthorHandler struct {
	AuthorUsecase models.AuthorUsecase
}

func NewAuthorHandler(e *echo.Echo, authorUsecase models.AuthorUsecase) {
	handler := &AuthorHandler{
		AuthorUsecase: authorUsecase,
	}
	e.GET("/authors", handler.getAll)
	e.POST("/authors", handler.Create)
	e.PUT("/authors", handler.Update)
	e.DELETE("/authors/:id", handler.Delete)
}

func (ch *AuthorHandler) getAll(c echo.Context) error {
	listArr, err := ch.AuthorUsecase.Get()
	if err != nil {
		log.Fatal("error while getting article: ", err)
	}
	return c.JSON(http.StatusOK, listArr)
}

func (ch *AuthorHandler) Create(c echo.Context) error {
	author := new(models.Author)
	if err := c.Bind(&author); err != nil {
		return err
	}
	err := ch.AuthorUsecase.Create(author)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Author has been successfully added")
}

func (ch *AuthorHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Please specify the id parameter")
	}
	id := int64(idP)
	err = ch.AuthorUsecase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (ch *AuthorHandler) Update(c echo.Context) error {
	author := new(models.Author)
	if err := c.Bind(&author); err != nil {
		return err
	}
	err := ch.AuthorUsecase.Update(author)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "Author has been successfully updated")
}
