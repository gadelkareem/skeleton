package tests

import (
    "backend/di"
    "backend/services"
    "github.com/brianvoe/gofakeit/v4"
    "github.com/danilopolani/gocialite/structs"
    "golang.org/x/oauth2"
    "gopkg.in/danilopolani/gocialite.v1"
)

var GocialiteMockUser *structs.User

type GocialiteMock struct {
}

func initSocialAuthService(c *di.Container) {
    c.SocialAuthService = services.NewSocialAuthService(c.UserService, c.JWTService, new(GocialiteMock))
}

func (g *GocialiteMock) New() *gocialite.Gocial {
    return gocialite.NewDispatcher().New()
}

func (g *GocialiteMock) Handle(state, code string) (*structs.User, *oauth2.Token, error) {
    GocialiteMockUser = &structs.User{
        FirstName: gofakeit.FirstName(),
        LastName:  gofakeit.LastName(),
        Email:     gofakeit.Email(),
    }
    return GocialiteMockUser, nil, nil
}
