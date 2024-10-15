package controller

import (
	"github.com/gin-gonic/gin"
	domain2 "go-backend-clean-architecture/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type TaskController struct {
	TaskUsecase domain2.TaskUsecase
}

func (tc *TaskController) Create(c *gin.Context) {
	var task domain2.Task

	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")
	task.ID = primitive.NewObjectID()

	task.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.TaskUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain2.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (u *TaskController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	tasks, err := u.TaskUsecase.FetchByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
