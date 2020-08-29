package services

import (
    "fmt"
    "reflect"
    "strings"
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

const ttl = 5 * time.Minute

func NewCacheService(c cachita.Cache, b bool) *CacheService {
    return &CacheService{c: c, debug: b}
}

func (s *CacheService) Get(i interface{}, p **paginator.Paginator, params ...interface{}) (id string, err error) {
    if kernel.App.DisableCache {
        return
    }
    if bs, k := i.(models.BaseInterface); k {
        params = append([]interface{}{s.modelType(bs)}, params...)
    }
    id, err = utils.Hash(params, p)
    if err != nil {
        s.log("Error hashing cache: %v", err)
        return
    }

    var w sync.WaitGroup
    w.Add(1)
    go func() {
        e := s.c.Get(id, i)
        if e != nil && !cachita.IsErrorOk(e) {
            s.log("Error retrieving cache: %v", e)
            err = e
        }
        w.Done()
    }()
    if p != nil {
        w.Add(1)
        go func() {
            e := s.c.Get(fmt.Sprintf("%s_paginator", id), p)
            if e != nil && !cachita.IsErrorOk(e) {
                s.log("Error retrieving cache: %v", e)
                err = e
            }
            w.Done()
        }()
    }
    w.Wait()

    return
}

func (s *CacheService) Put(i interface{}, id string, p *paginator.Paginator, tags ...string) (err error) {
    if kernel.App.DisableCache || id == "" {
        return
    }

    tags = s.getTags(i, p, tags...)

    var w sync.WaitGroup
    if p != nil {
        w.Add(1)
        go func() {
            p.Models = nil
            pID := fmt.Sprintf("%s_paginator", id)
            e := s.c.Put(pID, p, ttl)
            if e != nil {
                s.log("Error saving paginator cache: %v", e)
                err = e
                return
            }
            e = s.tag(pID, tags...)
            if e != nil {
                err = e
            }
        }()
    }
    go func() {
        w.Add(1)
        e := s.c.Put(id, i, ttl)
        if e != nil {
            s.log("Error saving cache: %v", e)
            err = e
            return
        }
        e = s.tag(id, tags...)
        if e != nil {
            err = e
        }
    }()
    w.Wait()
    return
}

func (s *CacheService) getTags(i interface{}, p *paginator.Paginator, tags ...string) (t []string) {
    t = tags
    if bs, k := i.(models.BaseInterface); k {
        t = append(t, s.modelTags(bs)...)
    }
    if p != nil {
        for _, m := range p.Models {
            if bs, k := m.(models.BaseInterface); k {
                t = append(t, s.modelTags(bs)...)
            }
        }
    }
    return
}

func (s *CacheService) tag(id string, tags ...string) (err error) {
    err = s.c.Tag(id, tags...)
    if err != nil {
        s.log("Error saving cache tags: %v", err)
    }
    return
}

func (s *CacheService) modelTags(m models.BaseInterface) []string {
    t := strings.TrimPrefix(reflect.TypeOf(m).String(), "*")
    return []string{t, fmt.Sprintf("%s_%s", t, m.GetID())}
}

func (s *CacheService) modelType(m models.BaseInterface) string {
    return strings.TrimPrefix(reflect.TypeOf(m).String(), "*")
}

func (s *CacheService) InvalidateModel(m models.BaseInterface) (err error) {
    err = s.Invalidate(s.modelTags(m)...)
    if err != nil {
        s.log("Error invalidating cache: %v", err)
    }
    return
}

func (s *CacheService) Invalidate(tags ...string) (err error) {
    if len(tags) == 0 {
        return nil
    }
    err = s.c.InvalidateTags(tags...)
    if err != nil {
        s.log("Error invalidating cache tags: %v", err)
    }
    return
}

func (s *CacheService) log(f interface{}, v ...interface{}) {
    if s.debug {
        logs.Error(f, v...)
    }
    return
}
