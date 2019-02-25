package gredis

import (
	"admin/pkg/logging"
	"admin/pkg/setting"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:setting.RedisSetting.MaxIdle,
		MaxActive:setting.RedisSetting.MaxActive,
		IdleTimeout:setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn,error){
			c,err := redis.Dial("tcp",setting.RedisSetting.Host)
			if err != nil {
				logging.Error(err)
				return nil,err
			}

			if setting.RedisSetting.Password != "" {
				if _,err := c.Do("AUTH",setting.RedisSetting.Password); err != nil {
					logging.Error(err)
					return nil,err
				}
			}
			return c,nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_,err := c.Do("PING")
			return err
		},
	}
	return nil
}

func Set(key string,data interface{},time int)(bool,error){
	conn := RedisConn.Get()
	defer conn.Close()

	value,err := json.Marshal(data)
	if err != nil {
		logging.Error(err)
		return false , err
	}

	reply,err := redis.Bool(conn.Do("SET",key,value))
	conn.Do("EXPIRE", key, time)

	return reply, err
}
func Get(key string) ([]byte,error){
	conn := RedisConn.Get()
	defer conn.Close()

	reply,err := redis.Bytes(conn.Do("GET",key))
	if err != nil {
		logging.Error(err)
		return nil,err
	}

	return reply,err
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists,err := redis.Bool(conn.Do("EXIST",key))
	if err != nil{
		logging.Error(err)
		return false
	}
	return exists
}

func Delete(key string) (bool,error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply,err := redis.Bool(conn.Do("DEL",key))
	if err != nil{
		logging.Error(err)
		return false,err
	}
	return reply,err
}

func LikeDeletes(key string) error{
	conn := RedisConn.Get()
	defer conn.Close()

	keys,err := redis.Strings(conn.Do("KEYS","*"+key+"*"))
	if err != nil {
		logging.Error(err)
		return err
	}

	for _,key := range keys {
		_,err := Delete(key)
		if err != nil {
			logging.Error(err)
			return err
		}
	}
	return nil
}
//key	键名
//data	键值
//ince	递增数
//time	过期时间
func ZSet(key string,data interface{},ince,time int) (bool,error){
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bool(conn.Do("ZADD", key, "INCR", ince, data))
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	conn.Do("EXPIRE", key, time)

	return reply,err
}
//RedisConn.Get()	在连接池中获取一个活跃连接
//conn.Do(commandName string, args ...interface{})	向 Redis 服务器发送命令并返回收到的答复
//redis.Bool(reply interface{}, err error)		将命令返回转为布尔值
//redis.Bytes(reply interface{}, err error)		将命令返回转为 Bytes
//redis.Strings(reply interface{}, err error)	将命令返回转为 []string
//redis.String(reply interface{}, err error)	将命令返回转为 string
//redis.Float64s(reply interface{}, err error)	将命令返回转为 []float64
//redis.Float64(reply interface{}, err error)	将命令返回转为 float64
//redis.Ints(reply interface{}, err error)		将命令返回转为 []Ints
//redis.Int(reply interface{}, err error)		将命令返回转为 Int
