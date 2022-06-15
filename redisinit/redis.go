package redisinit

import (
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type Redisobj struct {
	conn redis.Conn
}

// 生成redis连接
func (r *Redisobj) RedisConn(address, port, passwdstr string) error {
	var err error
	fmt.Println("开始连接redis数据库")
	url := fmt.Sprintf("%s:%s", address, port)
	db := redis.DialDatabase(0)
	passwd := redis.DialPassword(passwdstr)
	r.conn, err = redis.Dial("tcp", url, db, passwd)

	if err != nil {
		errors.New("连接redis失败")
		fmt.Printf("连接redis失败:%v\n", err)
		return err
	}
	fmt.Println("redis数据库连接成功")
	return err
}


// redis插入数据
func (r *Redisobj) Redisdo(method string, key, msg interface{}) {

	_, err := r.conn.Do(method, key, msg)
	if err != nil {
		errors.New("插入数据失败")
		fmt.Printf("插入数据失败:%v\n", err)
	}

}

// redis获取文件存在时的为偏移量
func (r *Redisobj) RRedisget(key string) (string, error) {
	name, err := redis.String(r.conn.Do("GET", key))
	if err != nil {
		fmt.Printf("查询位偏移量失败:%v\n", err)
		errors.New("查询位偏移量失败")
		return name, err
	}

	return name, err
}
