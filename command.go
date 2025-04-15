// System for easy command insertion and removal
package main

type command struct {
	Name string
	Args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

// This method registers a new handler function for a command name.
func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

// This method runs a given command with the provided state if it exists.
func (c *commands) run(s *state, cmd command) error {
	if cmd.Name == "login" {
		handlerLogin(s, cmd)
	} else if cmd.Name == "register" {
		handlerRegister(s, cmd)
	} else if cmd.Name == "reset" {
		handlerReset(s, cmd)
	} else if cmd.Name == "users" {
		handlerGetUsers(s, cmd)
	}
	return nil
}
