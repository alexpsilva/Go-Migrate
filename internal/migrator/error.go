package migrator

import "fmt"

type InvalidDbStringError struct{ Db_string string }

func (e InvalidDbStringError) Error() string {
	return fmt.Sprintf("Invalid database string '%v' provided", e.Db_string)
}

type InvalidMigrationsPathError struct{ Migrations_path string }

func (e InvalidMigrationsPathError) Error() string {
	return fmt.Sprintf("Invalid migrations path '%v' provided", e.Migrations_path)
}
