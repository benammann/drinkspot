package api_version

var Schema = `

	extend type Query {
		# Returns the current api version
		getApiVersion(): ApiVersion
	}

	# Represents the api version
	type ApiVersion {
			
		# Represents the app name
		app_name: String

		# Represents the app version
		app_version: String

	}

`
