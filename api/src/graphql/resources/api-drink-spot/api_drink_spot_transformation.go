package api_drink_spot

import "github.com/benammann/drinkspot-core/api/app/model"

func DrinkSpotToResolver(drinkSpot *model.DrinkSpot) *DrinkSpotResolver {

	var id32 = int32(drinkSpot.Model.ID)

	return &DrinkSpotResolver{
		DrinkSpot{
			Id:          &id32,
			Name:        &drinkSpot.Name,
			Description: &drinkSpot.Description,
			Latitude:    &drinkSpot.Latitude,
			Longitude:   &drinkSpot.Longitude,
			Quality:     &drinkSpot.Quality,
			UpVotes:     &drinkSpot.UpVotes,
			DownVotes:   &drinkSpot.DownVotes,
		},
	}
}
