package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remaingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumb := userTickets > 0 && userTickets <= remaingTickets
	return isValidName, isValidEmail, isValidTicketNumb
}
