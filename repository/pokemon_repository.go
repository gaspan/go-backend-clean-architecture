package repository

import (
	"context"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type pokemonRepository struct {
	database   mongo.Database
	collection string
}

func NewPokemonRepository(db mongo.Database, collection string) domain.PokemonRepository {
	return &pokemonRepository{
		database:   db,
		collection: collection,
	}
}

func (p *pokemonRepository) Create(c context.Context, pokemon *domain.PokemonCollection) error {
	collection := p.database.Collection(p.collection)

	_, err := collection.InsertOne(c, pokemon)

	return err
}

func (ur *pokemonRepository) ListMyPokemon(c context.Context) ([]domain.PokemonCollection, error) {
	collection := ur.database.Collection(ur.collection)

	cursor, err := collection.Find(c, bson.M{"is_release": 0})

	if err != nil {
		return nil, err
	}

	var pokemons []domain.PokemonCollection

	err = cursor.All(c, &pokemons)
	if pokemons == nil {
		return []domain.PokemonCollection{}, err
	}

	return pokemons, err
}

func (db *pokemonRepository) ReleasePokemon(c context.Context, name string) error {
	filter := bson.M{"name": name}
	update := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "is_release", Value: 1},
		},
	}}

	collection := db.database.Collection(db.collection)

	_, err := collection.UpdateMany(
		c,
		filter,
		update,
	)
	if err != nil {
		return err
	}
	return nil
}

func (db *pokemonRepository) RenamePokemon(c context.Context, name string, nickName string) error {
	filter := bson.M{"name": name}
	update := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "nick_name", Value: nickName},
		},
	}}

	collection := db.database.Collection(db.collection)

	_, err := collection.UpdateOne(
		c,
		filter,
		update,
	)
	if err != nil {
		return err
	}
	return nil
}
