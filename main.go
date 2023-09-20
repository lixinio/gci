package main

import (
	"os"

	"github.com/daixiang0/gci/cmd/gci"
)

var version = "v0.11.2r2"

func main() {
	e := gci.NewExecutor(version)

	err := e.Execute()
	if err != nil {
		os.Exit(1)
	}
}
