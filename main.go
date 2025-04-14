package main

import (
	"database/sql"
	"gator/internal/config"
	"gator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	loadConfigFromFile, err := config.Read()
	if err != nil {
		log.Print(err)
	}

	connStr := loadConfigFromFile.Db_url
	db, err := sql.Open("postgres", connStr)

	dbQuries := database.New(db)
	args := os.Args
	if len(args) <= 2 {
		os.Exit(1)
	}

	s := state{&loadConfigFromFile, dbQuries}
	cmds := commands{handlers: make(map[string]func(*state, command) error)}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	cmds.run(&s, command{args[1], args[2:]})
}
