package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	output := strings.Repeat("*", numStarsPerLine)
	output = output + "\n" + welcomeMsg
	output = output + "\n" + strings.Repeat("*", numStarsPerLine)
	return output
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	output := strings.ReplaceAll(oldMsg, "*", "")
	output = strings.Trim(output, "\n")
	output = strings.Trim(output, " ")
	return output
}
