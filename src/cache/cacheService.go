package cache

import (
	"KaduHod/muscles_api/src/core"
	"KaduHod/muscles_api/src/database"
	"KaduHod/muscles_api/src/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheService struct {
    Redis *database.Redis
}
func (self *CacheService) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cache, err := self.GetCacheFromRoute(r)
        if err != nil && err.Error() != "Cache miss" {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        if cache != nil && err == nil {
            w.Header().Set("Content-Type", "application/json")
            etag := utils.GenerateEtag(cache)
            if r.Header.Get("If-None-Match") == etag {
                w.WriteHeader(http.StatusNotModified)
                return
            }
            w.Header().Set("ETag", etag)
            w.WriteHeader(http.StatusOK)
            w.Write(cache)
            return
        }
        fmt.Println("Cache miss")
        next.ServeHTTP(w, r)
    })
}
func (self *CacheService) get(key string, dest interface{}) error {
    data, err := self.Redis.Conn.Get(context.Background(), key).Result()
    if err != nil && err.Error() == string(redis.Nil) {
        return errors.New("Cache miss")
    }
    if err != nil {
        return err
    }
    return json.Unmarshal([]byte(data), dest)
}
func (self *CacheService) set(key string, value interface{}, expiration time.Duration) error {
    data, err := json.Marshal(value)
    if err != nil {
        fmt.Println(err)
        return err
    }
    if err := self.Redis.Conn.Set(context.Background(), key, data, expiration).Err(); err != nil {
        fmt.Println(err)
        return err
    }
    return nil
}
func (self *CacheService) GetCacheFromRoute(r *http.Request) ([]byte, error) {
    key, err := self.PrepareKeyFromRoute(r)
    if err != nil {
        return nil, err
    }
    var resource interface{}
    if err := self.get(key, &resource); err != nil {
        return nil, err
    }
    json, err := json.Marshal(resource)
    if err != nil {
        return []byte(""), err
    }
    return json, nil
}
func (self *CacheService) PrepareKeyFromRoute(r *http.Request) (string, error) {
    var key string
    key = r.URL.Path
    key = strings.ReplaceAll(key, "/", "_")
    args := make(map[string][]string)
    keysSlice := []string{}
    if r.Method == http.MethodPost {
        if err := r.ParseForm(); err != nil {
            fmt.Println(err)
            return "", err
        }
        for k, values := range r.Form {
            args[k] = values
            keysSlice = append(keysSlice, k)
        }
    }
    for k := range r.URL.Query() {
        args[k] = r.URL.Query()[k]
        keysSlice = append(keysSlice, k)
    }
    sort.Strings(keysSlice)
    for _, k := range keysSlice {
        newKey := "_" + k + "_" + strings.Join(args[k], "_")
        key += newKey
    }
    key = "route:" + key
    return key, nil
}
func (self *CacheService) SetCacheFromRoute(r *http.Request , response interface{}) error {
    key, err := self.PrepareKeyFromRoute(r)
    if err != nil {
        fmt.Println(err)
        return err
    }
    exp := time.Hour * 24 * 7
    if err := self.set(key, response, exp); err != nil {
        fmt.Println(err)
        return err
    }
    return nil
}
func (self *CacheService) SetTokensFromUser(login string, tokens []core.UserAPIToken) error {
    key := "tokens:" + login
    exp := time.Hour * 24
    if err := self.set(key, tokens, exp); err != nil {
        return err
    }
    return nil
}
func (self *CacheService) GetTokensFromUser(login string) ([]core.UserAPIToken, error) {
    var tokens []core.UserAPIToken
    key := "tokens:" + login
    if err := self.get(key, &tokens); err != nil {
        if err.Error() == "Cache miss" {
            return tokens, nil
        }
        return nil, err
    }
    return tokens, nil
}
