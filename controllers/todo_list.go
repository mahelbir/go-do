package controllers

import (
	"github.com/gin-gonic/gin"
	"go-do/models"
	"go-do/services"
	"go-do/utils"
	"net/http"
)

func CreateTodoList(c *gin.Context) {
	details := todoListDetails(c)
	if details == nil {
		return
	}

	details.UserID = c.GetInt("userID")
	if details.Title == "" {
		c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Title is required", nil))
		return
	}

	todoList, err := services.TodoListSvc.Create(*details)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse())
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(true, "Todo list created", &todoList))
}

func UpdateTodoList(c *gin.Context) {
	details := todoListDetails(c)
	if details == nil {
		return
	}

	todoList := c.MustGet("todoList").(models.TodoList)
	if details.Title != "" {
		todoList.Title = details.Title
	}

	if _, err := services.TodoListSvc.Update(todoList); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse())
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(true, "Todo list updated", &todoList))
}

func DeleteTodoList(c *gin.Context) {
	todoList := c.MustGet("todoList").(models.TodoList)
	if err := services.TodoListSvc.Delete(todoList.ID); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse())
		return
	}

	c.JSON(http.StatusNoContent, utils.NewResponse(true, "Todo list deleted", nil))
}

func ListTodoList(c *gin.Context) {
	todoLists, err := services.TodoListSvc.ListByUserID(c.GetInt("userID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse())
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(true, "Todo lists retrieved", &todoLists))
}

func GetTodoList(c *gin.Context) {
	todoList := c.MustGet("todoList").(models.TodoList)
	c.JSON(http.StatusOK, utils.NewResponse(true, "Todo list retrieved", &todoList))
}

func todoListDetails(c *gin.Context) *models.TodoList {
	var todoListDetails models.TodoList
	if err := c.ShouldBindJSON(&todoListDetails); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Invalid request", nil))
		c.Abort()
		return nil
	}

	return &todoListDetails
}
