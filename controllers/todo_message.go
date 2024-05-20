package controllers

import (
	"github.com/gin-gonic/gin"
	"go-do/models"
	"go-do/services"
	"go-do/utils"
	"net/http"
)

func CreateTodoMessage(c *gin.Context) {
	details := todoMessageDetails(c)
	if details == nil {
		return
	}

	todoList := c.MustGet("todoList").(models.TodoList)
	details.TodoListID = todoList.ID
	if details.Content == "" {
		c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Content is required", nil))
		return
	}

	todoMessage, err := services.TodoMessageSvc.Create(*details)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse())
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(true, "Todo message created", todoMessage))
}

func UpdateTodoMessage(c *gin.Context) {
	details := todoMessageDetails(c)
	if details == nil {
		return
	}

	todoMessage := c.MustGet("todoMessage").(models.TodoMessage)
	details.IsCompleted = todoMessage.IsCompleted
	if details.Content != "" {
		todoMessage.Content = details.Content
	}

	if _, err := services.TodoMessageSvc.Update(todoMessage); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse())
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(true, "Todo message updated", todoMessage))
}

func DeleteTodoMessage(c *gin.Context) {
	todoMessage := c.MustGet("todoMessage").(models.TodoMessage)
	if err := services.TodoMessageSvc.Delete(todoMessage.ID); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse())
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(true, "Todo message deleted", nil))
}

func ListTodoMessageByTodoListID(c *gin.Context) {
	todoList := c.MustGet("todoList").(models.TodoList)
	todoMessages, err := services.TodoMessageSvc.ListByTodoListID(todoList.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse())
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(true, "Todo messages retrieved", todoMessages))
}

func GetTodoMessage(c *gin.Context) {
	todoMessage := c.MustGet("todoMessage").(models.TodoMessage)
	c.JSON(http.StatusOK, utils.NewResponse(true, "Todo message retrieved", todoMessage))
}

func todoMessageDetails(c *gin.Context) *models.TodoMessage {
	var details models.TodoMessage
	if err := c.ShouldBindJSON(&details); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(false, err.Error(), nil))
		return nil
	}
	return &details
}
