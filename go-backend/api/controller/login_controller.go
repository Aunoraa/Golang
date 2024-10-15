package controller

import (
	"github.com/gin-gonic/gin"
	"go-backend-clean-architecture/config"
	domain2 "go-backend-clean-architecture/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// @Summary Login user
// @Description Login user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param loginRequest body domain2.LoginRequest true "Login request"
// @Success 200 {object} domain2.LoginResponse
// @Failure 400 {object} domain2.ErrorResponse
// @Failure 404 {object} domain2.ErrorResponse
// @Failure 500 {object} domain2.ErrorResponse
// @Router /login [post]

type LoginController struct {
	LoginUsecase domain2.LoginUsecase
	Env          *config.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	var request domain2.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := lc.LoginUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain2.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain2.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := domain2.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
