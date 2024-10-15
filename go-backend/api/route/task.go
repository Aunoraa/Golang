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

func NewTaskRouter(env *configs.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
