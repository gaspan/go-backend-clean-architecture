package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/bootstrap"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/gin-gonic/gin"
)

type PokemonController struct {
	PokemonUsecase domain.PokemonUsecase
	Env            *bootstrap.Env
}

func (rtc *PokemonController) CatchPokemon(c *gin.Context) {

	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Float64()
	var hasil int
	if randNumber < 0.5 {
		hasil = 1
	} else {
		hasil = 0
	}

	if hasil == 1 {

		requestBody := domain.CathPokemonRequest{}
		err := c.BindJSON(&requestBody)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "Missing param",
				"data":    err,
			})
		}

		rawData := domain.PokemonCollection{
			Name:            requestBody.Name,
			ImgFrontDefault: requestBody.ImgFrontDefault,
			NickName:        "",
			IsRelease:       0,
		}
		err = rtc.PokemonUsecase.CatchPokemon(c, &rawData)
		fmt.Println("err:", err)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    500,
				"message": "Internal Server Error",
				"data":    err,
			})
		}

	}

	catchPokemonResponse := domain.CathPokemonResponse{
		IsCatched: hasil,
	}

	c.JSON(http.StatusOK, catchPokemonResponse)

}

func (rtc *PokemonController) GetListPokemon(c *gin.Context) {
	var request struct {
		Page int `form:"page,default=1"`
	}
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Missing param",
			"data":    err,
		})
	}

	pokemons, err := rtc.PokemonUsecase.GetListPokemon(c, request.Page)

	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Failed",
			"data":    err,
		})
	}

	if len(pokemons.Results) > 0 {

		for i, item := range pokemons.Results {
			detail := domain.DetailWithImage{
				Name:         item.Name,
				FrontDefault: "",
				BackDefault:  "",
			}
			detailAPI, err := rtc.PokemonUsecase.GetDetailPokemon(c, item.Name)
			if err == nil {
				detail = domain.DetailWithImage{
					Name:         item.Name,
					FrontDefault: detailAPI.Sprites.FrontDefault,
					BackDefault:  detailAPI.Sprites.BackDefault,
				}
			}
			pokemons.Results[i].Detail = detail
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    pokemons,
	})

}

func (rtc *PokemonController) DetailPokemon(c *gin.Context) {
	name := c.Param("name")

	DetailPokemon, err := rtc.PokemonUsecase.GetDetailPokemon(c, name)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Failed",
			"data":    err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    DetailPokemon,
	})

}

func (rtc *PokemonController) MyPokemonList(c *gin.Context) {

	MyPokemons, err := rtc.PokemonUsecase.GetMyPokemons(c)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    500,
			"message": "Failed",
			"data":    err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    MyPokemons,
	})

}

func (rtc *PokemonController) ReleasePokemon(c *gin.Context) {

	requestBody := domain.ReleasePokemonRequest{}
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(400, gin.H{
			"code":            400,
			"message":         "Missing param",
			"number":          0,
			"is_prime_number": false,
		})
	}

	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Int63()

	if randNumber > 1 && randNumber%2 != 0 {
		err = rtc.PokemonUsecase.ReleaseMyPokemons(c, requestBody.Name)
		if err != nil {
			c.JSON(400, gin.H{
				"code":            500,
				"message":         "Failed",
				"number":          0,
				"is_prime_number": false,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"code":            http.StatusOK,
			"message":         "Success",
			"number":          randNumber,
			"is_prime_number": true,
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":            400,
			"message":         "failed",
			"number":          randNumber,
			"is_prime_number": false,
		})
	}

}

func (rtc *PokemonController) RenamePokemon(c *gin.Context) {
	requestBody := domain.RenamePokemonRequest{}
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Missing param",
			"data":    "",
		})
	}
	err = rtc.PokemonUsecase.RenameMyPokemons(c, requestBody.Name, requestBody.GivenName)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "Internal Server Error",
			"data":    "",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    "",
	})
}
