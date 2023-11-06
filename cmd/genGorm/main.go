package main

import (
	"flag"

	"gorm.io/gen"

	"github.com/Flash-Pass/flash-pass-server/config"
	"github.com/Flash-Pass/flash-pass-server/db"
	"github.com/Flash-Pass/flash-pass-server/db/model"
)

var (
	username = flag.String("username", "root", "username")
	password = flag.String("password", "root", "password")
	host     = flag.String("host", "localhost", "host")
	port     = flag.Int("port", 3306, "port")
	database = flag.String("database", "flash_pass", "database")
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./db/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	db, err := db.InitMySQL(config.MySQLConfig{
		Username: *username,
		Password: *password,
		Host:     *host,
		Port:     *port,
		Database: *database,
	})
	if err != nil {
		panic(err)
	}

	g.UseDB(db)

	g.ApplyBasic(model.Card{}, model.User{})
	g.ApplyInterface(func(model.CardQueries) {}, model.Card{})

	g.Execute()
}
