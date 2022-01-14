package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	var res string = "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	return res
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var starsLine string = strings.Repeat("*", numStarsPerLine)
	return starsLine + "\n" + welcomeMsg + "\n" + starsLine
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var withoutStars string = strings.ReplaceAll(oldMsg, "*", "")
	return strings.TrimSpace(withoutStars)
}
