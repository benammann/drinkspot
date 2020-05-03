package graphql

import (
	"context"
	api_drink_spot "github.com/benammann/drinkspot-core/api/graphql/resources/api-drink-spot"
	"github.com/benammann/drinkspot-core/api/utility"
)

func (r *Resolver) CreateDrinkSpot(ctx context.Context, args *api_drink_spot.CreateDrinkSpotArgs) (*api_drink_spot.DrinkSpotResolver, error) {
	resolver, err := utility.GinRichQueryWithArgs(ctx, args, api_drink_spot.Mutation_CreateDrinkSpot)
	if err != nil {
		return nil, err
	} else {
		return resolver.(*api_drink_spot.DrinkSpotResolver), err
	}
}
