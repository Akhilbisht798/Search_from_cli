package main

import (
	"os"
	"strings"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	argsString := strings.Join(os.Args[1:], "+")
	search := "http://www.google.com/search?q=" + argsString
	open.Run(search)
}
