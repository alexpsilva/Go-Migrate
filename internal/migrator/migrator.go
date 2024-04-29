package migrator

import "fmt"

type Migrator struct {
	Db_url          string
	Migrations_path string
}

func (m *Migrator) Up() {
	fmt.Println("Migrating Up!")
}

func (m *Migrator) Down() {
	fmt.Println("Migrating Down!")
}
