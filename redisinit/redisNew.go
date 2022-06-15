package redisinit

import (
	"fmt"

	"github.com/go-redis/redis"
)

type Redisobjs struct {
	redisdb *redis.Client
}

// 初始化redis
func (r *Redisobjs) InitRedis(address, port, passwdstr string) (err error) {
	r.redisdb = redis.NewClient(&redis.Options{
		Addr: address + ":" + port, // 指定
		// Password: passwdstr,
		DB: 0, // redis一共16个库，指定其中一个库即可
	})
	_, err = r.redisdb.Ping().Result()
	if err != nil {
		fmt.Printf("redis数据库连接失败:%v\n", err)
	}
	return
}

// rpush写入数据
func (r *Redisobjs) RedisRPush(key string, msg interface{}) {
	r.redisdb.RPush(key, msg)
}

// set写入数据
func (r *Redisobjs) RedisSet(key string, msg int64) {
	// r.redisdb.Set(key, msg, time.Minute*10)
	r.redisdb.Set(key, msg, 0) // 设置为0永不过期
}

// get查询数据
func (r *Redisobjs) RedisGet(key string) string {
	cmd := r.redisdb.Get(key)
	res := cmd.Val()
	return res
}
