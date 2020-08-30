package di

import (
    "net/http"
    "os"

    "backend/commands"
    "backend/kernel"
    "backend/limiter"
    "backend/models"
    "backend/queue"
    "backend/queue/workers"
    "backend/rbac"
    "backend/services"
    "github.com/gadelkareem/cachita"
    h "github.com/gadelkareem/go-helpers"
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
    c.init()
    c.initCommands()
    c.initWorkers()

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

func (c *Container) init() {
    c.Cache = kernel.Cache()
    c.initQue()
    c.EmailService = services.NewEmailService(kernel.SMTPDialer(), nil, c.QueManager)
    c.SMSService = services.NewSMSService(&http.Client{}, c.QueManager)
    c.RateLimiter = limiter.NewRateLimiter(cachita.Memory(), nil)
    c.RBAC = rbac.New(kernel.IsDev())
    c.commonInit()
    c.SocialAuthService = services.NewSocialAuthService(c.UserService, c.JWTService, gocialite.NewDispatcher())
}

func (c *Container) initCommands() {
    kernel.Commands = map[string]kernel.Command{
        "migrate": commands.NewMigrator(kernel.DB()),
        "admin":   commands.NewAdmin(c.UserService),
    }
}

func (c *Container) InitTest() {
    c.commonInit()
}

func (c *Container) initQue() {
    if len(os.Args) > 1 {
        return
    }
    qc, _, err := kernel.Que(10)
    h.PanicOnError(err)
    c.QueManager = queue.NewQueManager(qc)
}

func (c *Container) initWorkers() {
    if c.QueManager == nil {
        return
    }
    c.QueManager.AddWorker(
        workers.NewSendMail(c.EmailService),
        workers.NewSendSMS(c.SMSService),
    )
}
