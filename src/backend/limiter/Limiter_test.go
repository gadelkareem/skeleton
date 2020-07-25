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

func TestLimiter_IsReached(t *testing.T) {
    t.Parallel()
    m := cachita.Memory()
    g := limiter.New("10-s", m, 10, time.Second)
    f := func(ip string, result bool) {
        b, err := g.IsExceeded(ip)
        tests.FailOnErr(t, err)
        assert.Equal(t, result, b)
    }
    home := "127.0.0.1"
    for i := 0; i < 10; i++ {
        f(home, false)
        f(h.RandomString(100), false)
    }
    f(home, true)
}
