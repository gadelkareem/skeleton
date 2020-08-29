package di

import (
    "net/http"

    "backend/commands"
    "backend/kernel"
    "backend/limiter"
    "backend/models"
    "backend/queue"
    "backend/queue/workers"
    "backend/rbac"
    "backend/services"
    "github.com/gadelkareem/cachita"
    "gopkg.in/danilopolani/gocialite.v1"
)

type Container struct {
    DB                   *kernel.PgDb
    UserRepository       *models.UserRepository
    UserService          *services.UserService
    EmailService         *services.EmailService
    JWTService           *services.JWTService
    SocialAuthService    *services.SocialAuthService
    AuditLogRepository   *models.AuditLogRepository
    AuditLogService      *services.AuditLogService
    AuthenticatorService *services.AuthenticatorService
    SMSService           *services.SMSService
    Cache                cachita.Cache
    CacheService         *services.CacheService
    RateLimiter          *limiter.RateLimiter
    RBAC                 *rbac.RBAC
    QueManager           *queue.QueManager
}

func InitContainer() *Container {
    c := new(Container)
    c.Init()
    c.InitCommands()
    c.InitWorkers()

    return c
}

func (c *Container) commonInit() {
    c.DB = kernel.NewDB()
    c.CacheService = services.NewCacheService(c.Cache, true)

    c.UserRepository = models.NewUserRepository(c.DB, 0)
    c.UserService = services.NewUserService(c.UserRepository, c.EmailService, c.SMSService, c.RBAC, c.CacheService)

    c.JWTService = services.NewJWTService(kernel.App.Config.String("hmacKey"), c.UserService)
    c.AuthenticatorService = services.NewAuthenticatorService(c.UserService)

    c.AuditLogRepository = models.NewAuditLogRepository(c.DB, 0)
    c.AuditLogService = services.NewAuditLogService(c.AuditLogRepository)
}

func (c *Container) Init() {
    c.Cache = kernel.Cache()
    qc, _ := kernel.Que(10)
    c.QueManager = queue.NewQueManager(qc)
    c.EmailService = services.NewEmailService(kernel.SMTPDialer(), nil, c.QueManager)
    c.SMSService = services.NewSMSService(&http.Client{})
    c.RateLimiter = limiter.NewRateLimiter(cachita.Memory(), nil)
    c.RBAC = rbac.New(kernel.IsDev())
    c.commonInit()
    c.SocialAuthService = services.NewSocialAuthService(c.UserService, c.JWTService, gocialite.NewDispatcher())
}

func (c *Container) InitCommands() {
    kernel.Commands = map[string]kernel.Command{
        "migrate": commands.NewMigrator(kernel.DB()),
        "admin":   commands.NewAdmin(c.UserService),
    }
}

func (c *Container) InitTest() {
    c.commonInit()
}

func (c *Container) InitWorkers() {
    c.QueManager.AddWorker(
        workers.NewSendMail(c.EmailService),
    )
    c.QueManager.StartWorkers()
}
