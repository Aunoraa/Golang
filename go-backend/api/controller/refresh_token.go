package controller

import (
	"github.com/gin-gonic/gin"
	"go-backend-clean-architecture/config"
	domain2 "go-backend-clean-architecture/internal/domain"
	"net/http"
)

type RefreshTokenController struct {
	RefreshTokenUsecase domain2.RefreshTokenUsecase
	Env                 *config.Env
}

func (rtc *RefreshTokenController) RefreshToken(c *gin.Context) {
	var request domain2.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	id, err := rtc.RefreshTokenUsecase.ExtractIDFromToken(request.RefreshToken, rtc.Env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain2.ErrorResponse{Message: "User not found"})
		return
	}

	user, err := rtc.RefreshTokenUsecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain2.ErrorResponse{Message: "User not found"})
		return
	}

	accessToken, err := rtc.RefreshTokenUsecase.CreateAccessToken(&user, rtc.Env.AccessTokenSecret, rtc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := rtc.RefreshTokenUsecase.CreateRefreshToken(&user, rtc.Env.RefreshTokenSecret, rtc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenResponse := domain2.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
