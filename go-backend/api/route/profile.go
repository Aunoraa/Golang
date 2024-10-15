package route

import (
	"github.com/gin-gonic/gin"
	"go-backend-clean-architecture/api/controller"
	"go-backend-clean-architecture/configs"
	"go-backend-clean-architecture/internal/domain"
	"go-backend-clean-architecture/internal/repository"
	"go-backend-clean-architecture/internal/usecase"
	"go-backend-clean-architecture/mongo"
	"time"
)

func NewProfileRouter(env *configs.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
