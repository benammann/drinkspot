package validation

// validates the password strength
func ValidatePasswordStrength(password string) bool {
	return len(password) > 5
}

// checks if the passwords match
func ValidatePasswordConfirmation(password string, passwordConfirmation string) bool {
	return password == passwordConfirmation
}
