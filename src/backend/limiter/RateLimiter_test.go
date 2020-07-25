package limiter_test

import (
    "testing"
    "time"

    "backend/limiter"
    "backend/tests"
    "github.com/gadelkareem/cachita"
    h "github.com/gadelkareem/go-helpers"
    "github.com/stretchr/testify/assert"
)

func TestRateLimiter_IsRateExceeded(t *testing.T) {
    t.Parallel()
    r := limiter.NewRateLimiter(cachita.Memory(), []limiter.Rate{
        {L: &limiter.Limiter{Name: "api_10_mcs", Limit: 10, TTL: 10 * time.Millisecond},
            Routes: [][]string{
                {"*", "/api/*"},
            },
        },
        {L: &limiter.Limiter{Name: "api_5_ns", Limit: 3, TTL: 2 * time.Millisecond},
            Routes: [][]string{
                {"*", "/api/v1/auth/token"},
            },
        },
    })

    f := func(_ip string, result bool, u string) {
        b, err := r.IsRateExceeded(nil, _ip, u, "POST")
        tests.FailOnErr(t, err)
        assert.Equal(t, result, b)
    }

    ip := h.RandomString(10)
    for i := 1; i <= 10; i++ {
        f(ip, false, "/api/any")
        f(h.RandomString(100), false, "/api/any")
    }
    f(ip, true, "/api/any")
    f(h.RandomString(100), false, "/api/any")
    time.Sleep(10 * time.Millisecond)
    f(ip, false, "/api/any")

    ip = h.RandomString(10)
    for i := 1; i <= 3; i++ {
        f(ip, false, "/api/v1/auth/token")
    }
    f(ip, true, "/api/v1/auth/token")
    time.Sleep(2 * time.Millisecond)
    f(ip, false, "/api/v1/auth/token")

}
