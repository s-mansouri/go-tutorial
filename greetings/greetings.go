package greetings

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello there, %s", name)
}
