package cli

import "fmt"

type InvalidCommandError struct {
	name string
}

func (e InvalidCommandError) Error() string {
	return fmt.Sprintf(
		"Invalid command '%v'. Must be a function taking either a string slice or no parameters at all",
		e.name,
	)
}
