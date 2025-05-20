package commands

import (
	"fmt"

	"backend/internal"
	"backend/kernel"
	"backend/models"
	"backend/services"
	"github.com/brianvoe/gofakeit/v6"
	h "github.com/gadelkareem/go-helpers"
)

type dev struct {
	s *services.UserService
}

func NewDev(s *services.UserService) kernel.Command {
	return &dev{s: s}
}

func (c *dev) Run(args []string) {
	if !kernel.IsDev() {
		fmt.Println("This command is only available in development mode.")
		return
	}
	kernel.App.DisableCache = true
	username := ""
	if len(args) > 1 {
		username = args[1]
	}
	if args[0] == "new-user" {
		pass := ""
		if len(args) > 2 {
			pass = args[2]
		}
		c.newUser(username, pass)
	} else if args[0] == "del-user" {
		c.delUser(username)
	} else {
		c.Help()
	}
}

func (c *dev) newUser(username, pass string) {
	u := models.NewUser()
	gofakeit.Struct(&u)
	if username != "" {
		u.Username = username
	}
	if pass != "" {
		u.Password = pass
	} else {
		pass = u.Password
	}
	u.Password = h.Md5(u.Password)
	u.CustomerID = ""

	err := c.s.SignUp(u)
	if err != nil {
		internalErr, k := err.(internal.Error)
		if k {
			fmt.Printf("\n\n\n%+v", internalErr.Object().Meta)
		} else {
			fmt.Printf("\n\n\n%+v", err)
		}
		return
	}
	u, err = c.s.FindUser(&models.User{Username: username}, false)
	h.PanicOnError(err)

	u.Activate()
	u.VerifyMobile()
	err = c.s.Save(u)
	h.PanicOnError(err)

	u.Password = pass
	fmt.Printf("\n\n\nUsername: %s\nPassword: %s\n\n", u.Username, pass)
}

func (c *dev) delUser(username string) {
	u, err := c.s.FindUser(&models.User{Username: username}, false)
	h.PanicOnError(err)
	err = c.s.DeleteUser(u.ID)
	h.PanicOnError(err)
	fmt.Println("\n\n\nUser deleted.")
}

func (c *dev) Help() {
	fmt.Printf(`
Usage: skeleton dev [options] ...

    Dev commands

Available options:
    new-user  generate a user
    del-user  delete a user
`)
}
