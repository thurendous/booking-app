package helper

import "strings"

// capitalize the first letter of the function name to make it public
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickest := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickest
}
