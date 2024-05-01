package main

import (
	"github.com/alexpsilva/go-migrate/internal/cli"
	"github.com/alexpsilva/go-migrate/internal/migrator"
)

func main() {
	migrator, err := migrator.New("123", "456")
	if err != nil {
		panic(err.Error())
	}

	cli := cli.New()

	cli.AddCommand("up", func(parameters []string) { migrator.Up() })
	cli.AddCommand("down", func(parameters []string) { migrator.Down() })

	cli.RunCommand("up", nil)
}
