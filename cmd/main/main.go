package main

import (
	"github.com/francobottoni/client-api/internal/database"
	"github.com/francobottoni/client-api/internal/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	migration "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

const (
	migrationsRootFolder     = "file://migrations"
	migrationsScriptsVersion = 1
)

func main() {
	_ = log.InitLogger()

	client := database.NewSqlClient("root@tcp(localhost:3306)/client_db")
	doMigrate(client, "client_db")

	mux := Routes()
	server := NewServer(mux)
	server.Run()
}

func doMigrate(client *database.MySqlClient, dbName string) {
	driver, _ := migration.WithInstance(client.DB, &migration.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationsRootFolder,
		dbName,
		driver,
	)

	if err != nil {
		log.Log().Error(err.Error())
		return
	}

	current, _, _ := m.Version()
	log.Log().Infof("current migration version in %d", current)
	err = m.Migrate(migrationsScriptsVersion)

	if err != nil && err.Error() == "no change" {
		log.Log().Info("no migration needed")
	}
}
