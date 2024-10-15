package controller

import (
	domain2 "github.com/amitshekhariitbhu/go-backend-clean-architecture/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
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
