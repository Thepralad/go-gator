package main

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"log"
	"math/rand"
	"os"
	"time"
)

// Login Handler
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("the args is empty :(")
	}
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)
	_ = cancel

	name := sql.NullString{cmd.Args[0], true}
	_, err := s.db.GetUser(ctxWithTimeout, name)
	if err != nil {
		log.Panic("user does not exists")
		return fmt.Errorf("user does not not exists")
	}

	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("not being able to switch user :(")
	}

	fmt.Println("User Switched Successfully")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("the args is empty :(")
	}

	name := sql.NullString{
		String: cmd.Args[0],
		Valid:  true,
	}
	id := int32(rand.Intn(100))
	args := database.CreateUserParams{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)
	_ = cancel
	_, err := s.db.GetUser(ctxWithTimeout, name)
	if err == nil {
		os.Exit(1)
	}
	user, err := s.db.CreateUser(ctxWithTimeout, args)
	_ = user
	if err != nil {
		return fmt.Errorf("not being able to switch user :(")
	}

	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("not being able to switch user :(")
	}
	return nil
}
