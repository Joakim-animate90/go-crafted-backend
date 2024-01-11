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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}