package limiter

import (
    "net"
    "strings"
    "sync"
    "time"

    "backend/kernel"
    "backend/models"
    "backend/utils"
    "github.com/gadelkareem/cachita"
)

type RateLimiter struct {
    botIpCacheMu sync.RWMutex
    BotIpCache   map[string]bool
    rates        []Rate
}

type Rate struct {
    L      *Limiter
    Routes [][]string
}

var config = []Rate{
    {L: &Limiter{Name: "api_4_s", Limit: 4, TTL: time.Second},
        Routes: [][]string{
            {"*", "/api/*"},
        },
    },
    {L: &Limiter{Name: "api_5_m", Limit: 5, TTL: time.Minute},
        Routes: [][]string{
            {"*", "/api/v1/auth/token"},
            {"*", "/api/v1/auth/callback"},
        },
    },
    {L: &Limiter{Name: "api_1_5m", Limit: 1, TTL: 5 * time.Minute},
        Routes: [][]string{
            {"*", "/api/v1/users/forgot-password"},
            {"*", "/api/v1/users/{userId}/send-verify-sms"},
            {"*", "/api/v1/common/contact"},
        },
    },
}

func NewRateLimiter(store cachita.Cache, cfg []Rate) *RateLimiter {
    if cfg != nil {
        config = cfg
    }
    r := &RateLimiter{rates: config, BotIpCache: make(map[string]bool)}
    for _, c := range config {
        c.L.Store = store
    }
    r.rates = config
    if !kernel.IsDev() {
        for _, i := range kernel.TrustedIPs {
            r.BotIpCache[i] = true
        }
    }
    return r
}

func (r *RateLimiter) IsRateExceeded(u *models.User, ip, uri, method string) (b bool, err error) {
    b, err = r.isExceeded(u, ip, uri, method)
    if err != nil {
        return
    }
    if b {
        b = !r.isTrustedBot(ip)
    }
    return
}

func (r *RateLimiter) isExceeded(u *models.User, ip, route, method string) (bool, error) {
    for _, rate := range r.rates {
        for _, rt := range rate.Routes {
            if utils.HasMethod(rt[0], method) && utils.HasRoute(u, rt[1], route) {
                b, err := rate.L.IsExceeded(ip)
                if err != nil || b {
                    return b, err
                }
            }
        }
    }
    return false, nil
}

func (r *RateLimiter) isTrustedBot(ip string) bool {
    r.botIpCacheMu.RLock()
    if _, exists := r.BotIpCache[ip]; exists {
        r.botIpCacheMu.RUnlock()
        return true
    }
    r.botIpCacheMu.RUnlock()
    addr, err := net.LookupAddr(ip)
    if err != nil {
        // logs.Error("Error getting host for ip %s Error: %v", ip, err)
        return false
    }

    ds := []string{"google.com.", "googlebot.com.", "msn.com.", "baidu.com.", "semrush.com.", "yahoo.com."}
    for _, i := range addr {
        for _, d := range ds {
            if strings.HasSuffix(i, d) {
                ips, err := net.LookupHost(i)
                if err != nil {
                    // logs.Error("Error getting host for ip %s Error: %v", ip, err)
                    return false
                }
                for _, _ip := range ips {
                    if ip == _ip {
                        r.botIpCacheMu.Lock()
                        r.BotIpCache[ip] = true
                        r.botIpCacheMu.Unlock()
                        return true
                    }
                }
            }
        }
    }
    return false
}
