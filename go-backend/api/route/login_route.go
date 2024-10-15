package route

import (
	"go-backend-clean-architecture/internal/domain"

	"time"

	"go-backend-clean-architecture/api/controller"

	"go-backend-clean-architecture/config"

	"go-backend-clean-architecture/mongo"

	"go-backend-clean-architecture/repository"

	"go-backend-clean-architecture/usecase"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *config.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
