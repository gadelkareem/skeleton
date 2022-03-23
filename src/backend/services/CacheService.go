package services

import (
    "fmt"
    "reflect"
    "sync"
    "time"

    "backend/kernel"
    "backend/models"
    "backend/utils"
    "backend/utils/paginator"
    "github.com/astaxie/beego/logs"
    "github.com/gadelkareem/cachita"
)

type (
    CacheService struct {
        c     cachita.Cache
        debug bool
    }
)

const ttl = 10 * time.Hour

func NewCacheService(c cachita.Cache, b bool) *CacheService {
    return &CacheService{c: c, debug: b}
}

func (s *CacheService) Get(i interface{}, p **paginator.Paginator, params ...interface{}) (k string, err error) {
    if kernel.App.DisableCache {
        return
    }

    k, err = s.Key(i, p, params...)
    if err != nil {
        return
    }

    var w sync.WaitGroup
    w.Add(1)
    go func() {
        err := s.c.Get(k, i)
        if err != nil && !cachita.IsErrorOk(err) {
            s.log("Error retrieving cache: %v", err)
        }
        w.Done()
    }()
    if p != nil {
        w.Add(1)
        go func() {
            err := s.c.Get(s.paginatorKey(k), p)
            if err != nil && !cachita.IsErrorOk(err) {
                s.log("Error retrieving cache: %v", err)
            }
            w.Done()
        }()
    }
    w.Wait()

    return
}

func (s *CacheService) paginatorKey(k string) string {
    return fmt.Sprintf("%s_paginator", k)
}

func (s *CacheService) Put(i interface{}, k string, p *paginator.Paginator, tags ...string) (err error) {
    if kernel.App.DisableCache || k == "" {
        return
    }

    tags = s.getTags(i, p, tags...)

    var w sync.WaitGroup
    if p != nil {
        w.Add(1)
        go func() {
            defer func() { w.Done() }()
            p.Models = nil
            pk := s.paginatorKey(k)
            err := s.c.Put(pk, p, ttl)
            if err != nil {
                s.log("Error saving paginator cache: %v", err)
                return
            }
            err = s.tag(pk, tags...)
            if err != nil {
                return
            }
        }()
    }
    w.Add(1)
    go func() {
        defer func() { w.Done() }()
        err = s.c.Put(k, i, ttl)
        if err != nil {
            s.log("Error saving cache: %v", err)
            return
        }
        err = s.tag(k, tags...)
        if err != nil {
            return
        }
    }()
    w.Wait()
    return
}

func (s *CacheService) getTags(i interface{}, p *paginator.Paginator, tags ...string) (t []string) {
    t = append(tags, s.interfaceTags(i)...)
    if p != nil {
        for _, m := range p.Models {
            t = append(t, s.interfaceTags(m)...)
        }
    }
    return
}

func (s *CacheService) tag(k string, tags ...string) (err error) {
    err = s.c.Tag(k, tags...)
    if err != nil {
        s.log("Error saving cache tags: %v", err)
    }
    return
}

func (s *CacheService) interfaceTags(i interface{}) (t []string) {
    t = append(t, s.interfaceType(i))
    if bs, k := i.(models.BaseInterface); k {
        t = append(t, s.ModelTags(bs)...)
    }

    return
}

func (s *CacheService) ModelTags(m models.BaseInterface, tags ...string) []string {
    t := s.interfaceType(m)
    id := m.GetID()
    if id == "" {
        return append(tags, t)
    }
    return append(tags, fmt.Sprintf("%s_%s", t, id))
}

func (s *CacheService) interfaceType(i interface{}) string {
    t := reflect.TypeOf(i)
    for {
        switch t.Kind() {
        case reflect.Ptr, reflect.Slice, reflect.Map:
            t = t.Elem()
        case reflect.Interface:
            return ""
        default:
            return t.String()
        }
    }
}

func (s *CacheService) InvalidateModel(m models.BaseInterface, tags ...string) (err error) {
    tags = append(tags, s.ModelTags(m)...)
    err = s.InvalidateTags(tags...)
    if err != nil {
        s.log("Error invalidating cache: %v", err)
    }
    return
}

func (s *CacheService) InvalidateTags(tags ...string) (err error) {
    if len(tags) == 0 {
        return nil
    }
    err = s.c.InvalidateTags(tags...)
    if err != nil {
        s.log("Error invalidating cache tags: %v %+v", err, tags)
    }
    return
}

func (s *CacheService) Invalidate(k string) (err error) {
    err = s.c.Invalidate(k)
    if err != nil {
        s.log("Error invalidating cache key: %s %+v", k, err)
    }
    return
}

// func (s *CacheService) InvalidatePaginator(k string) (err error) {
//     err = s.c.InvalidateMulti(k, s.paginatorKey(k))
//     if err != nil {
//         s.log("Error invalidating cache key: %s %+v", k, err)
//     }
//
//     return
// }

func (s *CacheService) log(f interface{}, v ...interface{}) {
	if s.debug {
		logs.Error(f, v...)
	}
	return
}

func (s *CacheService) Key(i interface{}, p **paginator.Paginator, params ...interface{}) (k string, err error) {
    params = append(params, s.interfaceType(i))

    k, err = utils.Hash(params, p)
    if err != nil {
        s.log("Error hashing cache: %v", err)
    }
    return
}
