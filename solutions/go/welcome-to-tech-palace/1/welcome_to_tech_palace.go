package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    capitalName := strings.ToUpper(customer)
    return "Welcome to the Tech Palace, " + capitalName
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    starsPerLine := strings.Repeat("*", numStarsPerLine)
    return starsPerLine + "\n" + welcomeMsg + "\n" + starsPerLine
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	removeStars := strings.ReplaceAll(oldMsg, "*", "")
    return strings.TrimSpace(removeStars)
    
}
