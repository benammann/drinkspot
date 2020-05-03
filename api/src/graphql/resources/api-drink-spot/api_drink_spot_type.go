package api_drink_spot

type DrinkSpot struct {
	Id          *int32   `json:"id"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
	Quality     *string  `json:"quality"`
	UpVotes     *int32   `json:"up_votes"`
	DownVotes   *int32   `json:"down_votes"`
}

type DrinkSpotResolver struct {
	DrinkSpot
}

func (d *DrinkSpotResolver) Id() *int32 {
	return d.DrinkSpot.Id
}

func (d *DrinkSpotResolver) Name() *string {
	return d.DrinkSpot.Name
}

func (d *DrinkSpotResolver) Description() *string {
	return d.DrinkSpot.Description
}

func (d *DrinkSpotResolver) Latitude() *float64 {
	return d.DrinkSpot.Latitude
}

func (d *DrinkSpotResolver) Longitude() *float64 {
	return d.DrinkSpot.Longitude
}

func (d *DrinkSpotResolver) Quality() *string {
	return d.DrinkSpot.Quality
}

func (d *DrinkSpotResolver) UpVotes() *int32 {
	return d.DrinkSpot.UpVotes
}

func (d *DrinkSpotResolver) DownVotes() *int32 {
	return d.DrinkSpot.DownVotes
}
