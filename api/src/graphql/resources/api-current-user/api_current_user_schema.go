package api_current_user

var Schema = `

	extend type Query {
		# Returns the current api version
		getCurrentUser(): CurrentUser
	}

	# Represents the api version
	# You must be authorized to use this function
	type CurrentUser {

		# Represents the app name
		email_address: String

	}

`
