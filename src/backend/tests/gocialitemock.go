package tests

import (
    "backend/di"
    "backend/services"
    "github.com/brianvoe/gofakeit/v4"
    "github.com/gadelkareem/gocialite/v2"
    "github.com/gadelkareem/gocialite/v2/structs"
    "golang.org/x/oauth2"
)

var GocialiteMockUser *structs.User

type GocialiteMock struct {
    d *gocialite.Dispatcher
}

func initSocialAuthService(c *di.Container) {
    c.SocialAuthService = services.NewSocialAuthService(c.UserService, c.JWTService, &GocialiteMock{d: gocialite.NewDispatcher(c.Cache)})
}

func (g *GocialiteMock) New() *gocialite.Gocial {
    return g.d.New()
}

func (g *GocialiteMock) Handle(state, code string) (*structs.User, *oauth2.Token, error) {
    GocialiteMockUser = &structs.User{
        FirstName: gofakeit.FirstName(),
        LastName:  gofakeit.LastName(),
        Email:     gofakeit.Email(),
    }
    return GocialiteMockUser, nil, nil
}
