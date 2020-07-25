package commands

import (
    "fmt"

    "backend/kernel"
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
    if len(args) > 1 && args[0] == "make-admin" {
        c.makeAdmin(args[1])
    } else {
        c.Help()
    }
}

func (c *admin) makeAdmin(username string) {
    err := c.s.MakeAdmin(username)
    h.PanicOnError(err)
    fmt.Printf("User %s is now an admin\n", username)
}

func (c *admin) Help() {
    fmt.Printf(`
Usage: skeleton admin [options] ...

    Controls users

Available options:
    make-admin [username]  make a user an admin   
`)
}
