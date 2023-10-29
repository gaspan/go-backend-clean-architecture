package domain

import (
	"context"

	"github.com/gin-gonic/gin"
)

const (
	CollectionPokemon = "pokemons"
)

type CathPokemonRequest struct {
	Name            string `json:"name"`
	ImgFrontDefault string `json:"img_front_default"`
}

type ReleasePokemonRequest struct {
	Name string `json:"name"`
}

type RenamePokemonRequest struct {
	Name      string `json:"name"`
	GivenName string `json:"given_name"`
}

type CathPokemonResponse struct {
	IsCatched int `json:"is_catched"`
}

type PokemonCollection struct {
	Name            string `bson:"name"`
	ImgFrontDefault string `bson:"img_front_default"`
	NickName        string `bson:"nick_name"`
	IsRelease       int    `bson:"is_release"`
}

type Pokemon struct {
	Name   string          `bson:"name"`
	Url    string          `bson:"url"`
	Detail DetailWithImage `json:"detail"`
}

type DetailWithImage struct {
	Name         string
	FrontDefault string
	BackDefault  string
}

type PokemonListResponse struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Pokemon `json:"results"`
}

type PokemonDetailResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Sprites struct {
		BackDefault  string `json:"back_default"`
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
	Moves []struct {
		Move struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"move"`
	} `json:"moves"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	BaseExperience int `json:"base_experience"`
}

type PokemonUsecase interface {
	GetListPokemon(c *gin.Context, page int) (PokemonListResponse, error)
	GetDetailPokemon(c *gin.Context, name string) (PokemonDetailResponse, error)
	CatchPokemon(c context.Context, pokemon *PokemonCollection) error
	GetMyPokemons(c context.Context) (list []PokemonCollection, err error)
	ReleaseMyPokemons(c context.Context, name string) error
	RenameMyPokemons(c context.Context, name string, givenName string) error
}

type PokemonRepository interface {
	Create(c context.Context, pokemon *PokemonCollection) error
	ListMyPokemon(c context.Context) ([]PokemonCollection, error)
	ReleasePokemon(c context.Context, name string) error
	RenamePokemon(c context.Context, name string, nickName string) error
}
