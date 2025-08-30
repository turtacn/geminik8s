package main

import (
	"fmt"
	"os"

	"github.com/turtacn/geminik8s/internal/app/cli"
)

func main() {
	if err := cli.NewRootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

//Personal.AI order the ending
