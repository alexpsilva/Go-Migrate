package main

import "github.com/alexpsilva/go-migrate/internal/migrator"

func main() {
	migrator := migrator.Migrator{
		Db_url:          "123",
		Migrations_path: "456",
	}
	migrator.Up()
}
