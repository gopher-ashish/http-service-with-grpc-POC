package model

import "github.com/internal/connection"

func MigrateDB(db connection.ConnectionPool) {
	drop(db)
	migrate(db)
}

func drop(db connection.ConnectionPool) {
	db.DropTableIfExists(
		&User{},
	)
}

// migrate database tables (if changed).
func migrate(db connection.ConnectionPool) {
	db.AutoMigrate(
		&User{},
	)
}
