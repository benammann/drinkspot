package api_drink_spot

var Schema = `

	extend type Query {
		# Searches for nearby drinking spots
		searchDrinkingSpots(latitude: Float!, longitude: Float!, radius: Int): [DrinkSpot!]!
	}

	extend type Mutation {
		createDrinkSpot(
			name: String!
			description: String!
			Latitude: Float!
			Longitude: Float!
			Quality: DrinkSpotWaterQuality!
		): DrinkSpot
	}

	# Represents a DrinkSpot
	type DrinkSpot {

		id: Int
		name: String
		description: String
		latitude: Float
		longitude: Float
		quality: String
		up_votes: Int
		down_votes: Int

	}

	enum DrinkSpotWaterQuality {
		DRINKABLE
		NOT_DRINKABLE
		UNKNOWN
	}

`
