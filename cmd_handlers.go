package main

import (
	"fmt"
)

//Login Handler
func handlerLogin(s *state, cmd command) error{
	if len(cmd.Args) < 1{
		return fmt.Errorf("The args is empty :(")
	}
	
	err := s.cfg.SetUser(cmd.Args[0])
	if err != nil{
		return fmt.Errorf("Not being able to switch user :(")
	}

	fmt.Println("User Switched Successfully")
	return nil
}
