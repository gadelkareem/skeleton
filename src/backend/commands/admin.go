package commands

import (
	"fmt"

	"backend/kernel"
	"backend/models"
	"backend/services"
	h "github.com/gadelkareem/go-helpers"
)

type admin struct {
	s *services.UserService
}

func NewAdmin(s *services.UserService) kernel.Command {
	return &admin{s: s}
}

func (c *admin) Run(args []string) {
	if len(args) > 1 {
		switch args[0] {
		case "make-admin":
			c.makeAdmin(args[1])
			return
		case "activate-user":
			c.activateUser(args[1])
			return
		}
	} else {
		c.Help()
	}
}

func (c *admin) makeAdmin(username string) {
	err := c.s.MakeAdmin(username)
	h.PanicOnError(err)
	fmt.Printf("User %s is now an admin\n", username)
}

func (c *admin) activateUser(username string) {
	u, err := c.s.FindUser(&models.User{Username: username}, true)
	h.PanicOnError(err)
	u.Activate()
	err = c.s.Save(u, "active", "email_verify_hash")
	h.PanicOnError(err)
	fmt.Printf("User %s is now an active\n", username)
}

func (c *admin) Help() {
	fmt.Printf(`
Usage: skeleton admin [options] ...

    Controls users

Available options:
    make-admin [username]  make a user an admin
    activate-user [username]  activate a user account
`)
}
