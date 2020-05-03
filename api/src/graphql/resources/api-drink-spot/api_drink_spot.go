package api_drink_spot

import (
	"github.com/benammann/drinkspot-core/api/app/model"
	"github.com/benammann/drinkspot-core/api/database"
	"github.com/benammann/drinkspot-core/api/utility"
	"github.com/gin-gonic/gin"
)

type SearchDrinkingSpotArgs struct {
	Latitude  float64
	Longitude float64
	Radius    *int32
}

type CreateDrinkSpotArgs struct {
	Name        string
	Description string
	Latitude    float64
	Longitude   float64
	Quality     string
}

func Query_SearchDrinkingSpots(ctx *gin.Context, args interface{}) (interface{}, error) {

	var searchResults []*model.DrinkSpot

	var resolvers []*DrinkSpotResolver

	database.Current.Connection.Find(&searchResults)

	for _, value := range searchResults {
		resolvers = append(resolvers, DrinkSpotToResolver(value))
	}

	return resolvers, nil
}

func Mutation_CreateDrinkSpot(ctx *gin.Context, args interface{}) (interface{}, error) {

	_, err := utility.ExtractCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	arg := args.(*CreateDrinkSpotArgs)

	newDrinkSpot := &model.DrinkSpot{
		Name:        arg.Name,
		Description: arg.Description,
		Latitude:    arg.Latitude,
		Longitude:   arg.Longitude,
		Quality:     arg.Quality,
		UpVotes:     0,
		DownVotes:   0,
	}

	database.Current.Connection.Create(newDrinkSpot)

	return DrinkSpotToResolver(newDrinkSpot), nil
}
