package main

import(
	"log"
	"gator/internal/config"
	"os"
)

type state struct{
	cfg *config.Config
}

func main(){
	loadConfigFromFile, err := config.Read()	
	if err != nil{
		log.Print(err)
	}

	args := os.Args
	if len(args) <= 2 {
    	os.Exit(1)
	}
	s := state{&loadConfigFromFile}
	cmds := commands{handlers: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	cmds.run(&s,command{args[1], args[2:]})
}
