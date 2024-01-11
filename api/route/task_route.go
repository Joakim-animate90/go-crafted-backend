package route

import (
	"time"

	"github.com/Joakim-animate90/go-crafted-backend/api/controller"
	"github.com/Joakim-animate90/go-crafted-backend/bootstrap"
	"github.com/Joakim-animate90/go-crafted-backend/domain"
	"github.com/Joakim-animate90/go-crafted-backend/mongo"
	"github.com/Joakim-animate90/go-crafted-backend/repository"
	"github.com/Joakim-animate90/go-crafted-backend/usecase"
	"github.com/gin-gonic/gin"
)

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}