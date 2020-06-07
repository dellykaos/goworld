package goworld

import "fmt"

// Hello writes Hello to the world
func Hello(name string) string {
	if len(name) == 0 {
		name = "world"
	}

	return fmt.Sprintf("Hello, %s!", name)
}
