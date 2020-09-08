package kernel

import (
    "net/url"
    "time"

    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego/validation"

    "github.com/gadelkareem/cachita"
    "github.com/gadelkareem/go-helpers"
    "github.com/go-gomail/gomail"
)

func Cache() cachita.Cache {
    t := App.ConfigOrEnvVar("cacheType", "CACHE_TYPE")

    if t == "memory" {
        return cachita.Memory()
    }

    return redis()
}

func redis() cachita.Cache {
    var redis cachita.Cache
    u, err := url.Parse(App.ConfigOrEnvVar("redisAddress", "REDIS_URL"))
    h.PanicOnError(err)
    pass, _ := u.User.Password()
    u.User = url.UserPassword("", pass)

    err = h.Retry(func() (e error) {
        redis, e = cachita.NewRedisCache(2*time.Hour, App.Config.DefaultInt("redisPoolSize", 100), "skeleton", u.String())
        if e != nil {
            logs.Error("Redis connection error: %s", e)
            time.Sleep(5 * time.Second)
        }
        return e
    }, MaxInt)
    h.PanicOnError(err)
    return redis
}

func SMTPDialer() *gomail.Dialer {
    return gomail.NewDialer(App.ConfigOrEnvVar("smtp::server", "MAIL_HOST"),
        App.Config.DefaultInt("smtp::port", 0),
        App.Config.String("smtp::smtpUser"),
        App.Config.String("smtp::smtpPassword"),
    )
}

func Validation() *validation.Validation {
    validation.CanSkipFuncs["MinSize"] = struct{}{}
    validation.CanSkipFuncs["Match"] = struct{}{}
    validation.CanSkipFuncs["Numeric"] = struct{}{}
    validation.CanSkipFuncs["Length"] = struct{}{}
    return &validation.Validation{RequiredFirst: true}
}
