package limiter

import (
    "fmt"
    "time"

    "github.com/gadelkareem/cachita"
)

type Limiter struct {
    Name  string
    Store cachita.Cache
    TTL   time.Duration
    Limit int64
}

func New(name string, store cachita.Cache, limit int64, ttl time.Duration) *Limiter {
    return &Limiter{
        Name:  name,
        Store: store,
        Limit: limit,
        TTL:   ttl,
    }
}

func (l *Limiter) k(key string) string {
    return fmt.Sprintf("limiter:%s:%s", l.Name, key)
}

func (l *Limiter) IsExceeded(key string) (bool, error) {
    count, err := l.Store.Incr(l.k(key), l.TTL)
    if err != nil {
        return false, err
    }
    return count > l.Limit, nil
}
