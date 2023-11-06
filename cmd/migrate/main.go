package main

import (
	"flag"
	"github.com/Flash-Pass/flash-pass-server/db"
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"gorm.io/gen"
)

var (
	username = flag.String("username", "root", "username")
	password = flag.String("password", "root", "password")
	host     = flag.String("host", "localhost", "host")
	port     = flag.String("port", "3306", "port")
	database = flag.String("database", "flash_pass", "database")
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./db/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	db, err := db.InitMySQL(*username, *password, *host, *port, *database)
	if err != nil {
		panic(err)
	}

	g.UseDB(db)

	g.ApplyBasic(model.Card{})
	g.ApplyInterface(func(model.CardQueries) {}, model.Card{})

	g.Execute()
}
