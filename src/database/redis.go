package database

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
    Conn *redis.Client
}

func NewRedis() *Redis {
    redisConn, err := redisConn(context.Background())
    if err != nil {
        panic(err)
    }
    return &Redis{Conn: redisConn}
}

func redisConn(ctx context.Context) (*redis.Client, error) {
    redisHost := os.Getenv("REDIS_HOST")
    password := os.Getenv("REDIS_PASSWORD")
    redisPort := os.Getenv("REDIS_PORT")
    database, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
    if err != nil {
        panic(err)
    }
    conn := redis.NewClient(&redis.Options{
        Addr: redisHost + ":" + redisPort,
        Password: password,
        DB: database,
    })
    pong, err := conn.Ping(ctx).Result()
    if err != nil {
        fmt.Println(pong, err, "Redis :: Connection result")
        return conn, err
    }

    return conn, nil
}
