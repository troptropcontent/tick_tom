package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/troptropcontent/tick_tom/db"
)

func requireSubcommand() {
	if len(os.Args) < 2 {
		fmt.Println("You must provide a subcommand")
		os.Exit(1)
	}
}

func main() {
	flag.NewFlagSet("create", flag.ExitOnError)
	requireSubcommand()
	switch os.Args[1] {
	case "create":
		db.Create()
	default:
		fmt.Println("Unknown subcommand")
		os.Exit(1)
	}
}
