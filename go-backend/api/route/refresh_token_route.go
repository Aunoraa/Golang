package route

import (
	"github.com/gin-gonic/gin"
	"go-backend-clean-architecture/api/controller"
	"go-backend-clean-architecture/config"
	"go-backend-clean-architecture/internal/domain"
	"go-backend-clean-architecture/mongo"
	"go-backend-clean-architecture/repository"
	"go-backend-clean-architecture/usecase"
	"time"
)

func NewRefreshTokenRouter(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
