package http

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"test_crud/models"
)

type CommentHandler struct {
	CommentUsecase models.CommentUsecase
}

func NewCommentHandler(e *echo.Echo, commentUsecase models.CommentUsecase) {
	handler := &CommentHandler{
		CommentUsecase: commentUsecase,
	}
	e.GET("/comments", handler.getAll)
	e.POST("/comments", handler.Create)
	e.PUT("/comments", handler.Update)
	e.DELETE("/comments/:id", handler.Delete)
}

func (ch *CommentHandler) getAll(c echo.Context) error {
	commentList, err := ch.CommentUsecase.Get()
	if err != nil {
		log.Fatal("error while getting article: ", err)
	}
	return c.JSON(http.StatusOK, commentList)
}

func (ch *CommentHandler) Create(c echo.Context) error {
	comment := new(models.Comment)
	if err := c.Bind(&comment); err != nil {
		return err
	}
	err := ch.CommentUsecase.Create(comment)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Comment has been successfully added")
}

func (ch *CommentHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Please specify the id parameter")
	}
	id := int64(idP)
	err = ch.CommentUsecase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (ch *CommentHandler) Update(c echo.Context) error {
	comment := new(models.Comment)
	if err := c.Bind(&comment); err != nil {
		return err
	}
	err := ch.CommentUsecase.Update(comment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "Comment has been successfully updated")
}
