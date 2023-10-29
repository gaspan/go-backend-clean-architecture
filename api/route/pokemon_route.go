package route

import (
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/api/controller"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/bootstrap"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/repository"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/usecase"
	"github.com/gin-gonic/gin"
)

func NewPokemonRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPokemonRepository(db, domain.CollectionPokemon)
	rtc := &controller.PokemonController{
		PokemonUsecase: usecase.NewPokemonUsecase(ur, timeout),
		Env:            env,
	}
	group.POST("/pokemon/catch", rtc.CatchPokemon)
	group.GET("/pokemon/list", rtc.GetListPokemon)
	group.GET("/pokemon/detail/:name", rtc.DetailPokemon)
	group.GET("/pokemon/my", rtc.MyPokemonList)
	group.POST("/pokemon/release", rtc.ReleasePokemon)
	group.POST("/pokemon/rename", rtc.RenamePokemon)
}
