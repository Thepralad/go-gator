package main

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"math/rand"
	"os"
	"time"
)

// Login Handler
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("The args is empty :(")
	}

	err := s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Not being able to switch user :(")
	}

	fmt.Println("User Switched Successfully")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("The args is empty :(")
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
	ctxWithTimeout, _ := context.WithTimeout(ctx, 3*time.Second)

	_, err := s.db.GetUser(ctxWithTimeout, name)
	if err == nil {
		os.Exit(1)
	}
	user, err := s.db.CreateUser(ctxWithTimeout, args)
	_ = user
	if err != nil {
		return fmt.Errorf("not being able to switch user :(")
	}
	return nil
}
