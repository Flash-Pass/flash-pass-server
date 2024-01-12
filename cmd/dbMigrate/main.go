package main

import (
	"flag"

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

	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.User{}, &model.Card{}, &model.Plan{}, &model.BookCard{}, &model.Book{}, &model.Plan{}, &model.Task{},
		&model.TaskCardRecord{}, &model.TaskLog{},
	)

	db.Migrator()

	if err != nil {
		panic(err)
	}
}
