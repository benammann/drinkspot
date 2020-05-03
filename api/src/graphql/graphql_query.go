package graphql

import (
	"context"
	api_current_user "github.com/benammann/drinkspot-core/api/graphql/resources/api-current-user"
	api_drink_spot "github.com/benammann/drinkspot-core/api/graphql/resources/api-drink-spot"
	api_version "github.com/benammann/drinkspot-core/api/graphql/resources/api-version"
	"github.com/benammann/drinkspot-core/api/utility"
)

func (r *Resolver) GetApiVersion(ctx context.Context) (*api_version.ApiVersionResolver, error) {
	resolver, err := utility.GinRichQuery(ctx, api_version.Query_GetApiVersion)
	if err != nil {
		return nil, err
	} else {
		return resolver.(*api_version.ApiVersionResolver), err
	}
}

func (r *Resolver) GetCurrentUser(ctx context.Context) (*api_current_user.CurrentUserResolver, error) {
	resolver, err := utility.GinRichQuery(ctx, api_current_user.Query_CurrentUser)
	if err != nil {
		return nil, err
	} else {
		return resolver.(*api_current_user.CurrentUserResolver), err
	}
}

func (r *Resolver) SearchDrinkingSpots(ctx context.Context, args api_drink_spot.SearchDrinkingSpotArgs) ([]*api_drink_spot.DrinkSpotResolver, error) {
	resolver, err := utility.GinRichQueryWithArgs(ctx, args, api_drink_spot.Query_SearchDrinkingSpots)
	if err != nil {
		return nil, err
	} else {
		return resolver.([]*api_drink_spot.DrinkSpotResolver), err
	}
}
