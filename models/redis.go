package models

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisConn *redis.Pool
var RedisSetting = &Redis{}

// Setup Initialize the Redis instance
func InitRedis() error {
	RedisConn = &redis.Pool{
		MaxIdle:     RedisSetting.MaxIdle,
		MaxActive:   RedisSetting.MaxActive,
		IdleTimeout: RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			if RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}

// Set a key/value
func SetStr(key string, data string, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, data)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

// Set a key/value
func SetInt(key string, data int, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, data)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func GetStr(key string) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := conn.Do("GET", key)
	if err != nil {
		return "", err
	}
	data, _ := redis.String(reply, err)
	return data, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
