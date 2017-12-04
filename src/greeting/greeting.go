package greeting

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("Hello, %v.", name)
}
