package appinit

import (
	"log"

	"ExBot/internal/texts"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dsn string) {
	m, err := migrate.New("file://../migrations", dsn)
	if err != nil {
		log.Fatalf(texts.LogMigrationsNew, err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf(texts.LogMigrationsUp, err)
	}
	log.Println(texts.LogMigrationsDone)
}
