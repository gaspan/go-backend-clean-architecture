package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/helpers"
	"github.com/gin-gonic/gin"
)

type pokemonUsecase struct {
	pokemonRepository domain.PokemonRepository
	contextTimeout    time.Duration
}

func NewPokemonUsecase(pokemonRepo domain.PokemonRepository, timeout time.Duration) domain.PokemonUsecase {
	return &pokemonUsecase{
		pokemonRepository: pokemonRepo,
		contextTimeout:    timeout,
	}
}

func (rtu *pokemonUsecase) GetListPokemon(c *gin.Context, page int) (resp domain.PokemonListResponse, err error) {

	limit := 10
	offset := (page - 1) * limit

	list, err := helpers.Get("https://pokeapi.co/api/v2/pokemon?offset="+strconv.Itoa(offset)+"&limit="+strconv.Itoa(limit), nil)
	if err != nil {
		return resp, err
	}
	var res domain.PokemonListResponse
	json.Unmarshal(list, &res)

	data := res

	return data, nil
}

func (rtu *pokemonUsecase) GetDetailPokemon(c *gin.Context, name string) (resp domain.PokemonDetailResponse, err error) {
	detail, err := helpers.Get("https://pokeapi.co/api/v2/pokemon/"+name, nil)
	if err != nil {
		return resp, err
	}
	var res domain.PokemonDetailResponse
	json.Unmarshal(detail, &res)

	data := res

	return data, nil
}

func (p *pokemonUsecase) CatchPokemon(c context.Context, pokemon *domain.PokemonCollection) error {
	fmt.Println("masuk usecase")
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	fmt.Println("masuk usecase 2")
	return p.pokemonRepository.Create(ctx, pokemon)
}

func (p *pokemonUsecase) GetMyPokemons(c context.Context) ([]domain.PokemonCollection, error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	list, err := p.pokemonRepository.ListMyPokemon(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (p *pokemonUsecase) ReleaseMyPokemons(c context.Context, name string) error {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	err := p.pokemonRepository.ReleasePokemon(ctx, name)
	if err != nil {
		return err
	}
	return nil
}

func printFibonacciSeries(num int) int {
	a := 0
	b := 1
	c := b

	lastNumber := 0

	for true {
		c = b
		b = a + b
		if b >= num {
			fmt.Println()
			return lastNumber
			// break
		}
		a = c
		lastNumber = b
	}
	return 0
}

var number int

func (p *pokemonUsecase) RenameMyPokemons(c context.Context, name string, givenName string) error {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	number += 1
	lastCombineName := "-" + strconv.Itoa(printFibonacciSeries(number))
	combineName := givenName + lastCombineName
	err := p.pokemonRepository.RenamePokemon(ctx, name, combineName)
	if err != nil {
		return err
	}
	return nil
}
