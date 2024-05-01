package migrator

import "fmt"

type Migrator struct {
	Db_string       string
	Migrations_path string
}

func New(db_string string, migrations_path string) (*Migrator, error) {
	if len(db_string) == 0 {
		return nil, InvalidDbStringError{Db_string: db_string}
	}

	if len(migrations_path) == 0 {
		return nil, InvalidMigrationsPathError{Migrations_path: migrations_path}
	}

	return &Migrator{
		Db_string:       db_string,
		Migrations_path: migrations_path,
	}, nil
}

func (m *Migrator) Up() {
	fmt.Println("Migrating Up!")
}

func (m *Migrator) Down() {
	fmt.Println("Migrating Down!")
}
