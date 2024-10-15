package controller

import (
	"github.com/gin-gonic/gin"
	domain2 "go-backend-clean-architecture/internal/domain"
	"net/http"
)

type ProfileController struct {
	ProfileUsecase domain2.ProfileUsecase
}

func (pc *ProfileController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	profile, err := pc.ProfileUsecase.GetProfileByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
