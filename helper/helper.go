package helper

import "strings"

func ValidateUserInput(FirstName string, LastName string, Email string, Tickets int, RemainingTickets int) (bool, bool, bool) {
	isValidName := len(FirstName) >= 2 && len(LastName) >= 2
	isValidEmail := strings.Contains(Email, "@")
	isValidTicketNumber := Tickets > 0 && Tickets <= RemainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
