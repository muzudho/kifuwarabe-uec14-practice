package greetings

import "fmt"

// GetMessage returns a greeting for the named person.
func GetMessage(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
