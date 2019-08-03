package greeting

import (
	"fmt"
	"strings"
)

func filterName(name string) string{
	naming := ""
	switch name{
	case "":
		naming="my friend"
	default:
		naming=name
	}
	return naming
}

func normalizeName(name string) string{
	if name == strings.ToUpper(name){
		return "HELLO " + name + "!"
	}
	return fmt.Sprintf("Hello, %v.", name)
}

func Greet(name string) string {
	naming := filterName(name)
	return normalizeName(naming)
}
