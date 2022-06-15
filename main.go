package main

import (
	"errors"
	"fmt"
	"logcollect/conf"
	"logcollect/redisinit"
	"strconv"
	"strings"
	"sync"

	"logcollect/tail"
)

var m sync.WaitGroup

// var redisc redisinit.Redisobj
var redisc redisinit.Redisobjs
var rlock sync.RWMutex

// 获取位偏移量
func getOffset(key string) int64 {
	rlock.RLock()
	// offset, err := redisc.RRedisget(key)
	offset := redisc.RedisGet(key)

	rlock.RUnlock()
	if offset == "" {
		fmt.Println("当前位偏移量为空")
		offset = "0"
	}

	// 统一将offset字符串转成数字

	of, _ := strconv.Atoi(offset)
	return int64(of)
}

// 加载日志
func redlogs(logpath string, redisconn redisinit.Redisobjs) {
	defer m.Done()
	// 获取文件名
	fileNameSplit := strings.Split(logpath, "/")
	fileName := fileNameSplit[len(fileNameSplit)-1]

	// 文件名去掉后缀名
	fileSplit := strings.Split(fileName, ".")
	filePrefix := fileSplit[0]

	// 获取位偏移量
	getoffset := getOffset(filePrefix + "offset")
	// 启动太累
	tailsteam := tail.TailRead(logpath, getoffset)
	for {
		msg, _ := <-tailsteam.Lines
		fmt.Println(msg.Text)
		offset, err := tailsteam.Tell() // 获取写入文件当前的为偏移量
		if err != nil {
			fmt.Println("未获取到的最新的位偏移量")
			errors.New("未获取到的最新的位偏移量")
		} else {
			fmt.Println("当前位偏移量", offset)
			// 写入当前的位偏移量
			rlock.Lock()
			// redisconn.Redisdo("set", filePrefix+"offset", offset)
			redisconn.RedisSet(filePrefix+"offset", offset)
			rlock.Unlock()
		}
		// 写入数据
		rlock.Lock()
		// redisconn.Redisdo("rpush", fileName, msg.Text)
		redisconn.RedisRPush(fileName, msg.Text)
		rlock.Unlock()
	}
}

func main() {
	// 加载配置文件

	conifg, err := conf.Reloadconf()
	if err != nil {
		errors.New("获取ini配置文件失败")
	}

	// 初始化redis

	// err = redisc.RedisConn(conifg.Redis.Address, conifg.Redis.Port, conifg.Redis.Passwd)
	err = redisc.InitRedis(conifg.Redis.Address, conifg.Redis.Port, conifg.Redis.Passwd)

	if err != nil {
		errors.New("连接redis失败")
	}

	// 获取目录路径
	// logpaths := strings.Split(conifg.Tail.Logpath, ",")
	logpath := conifg.Tail.Logpath
	fileNames := strings.Split(conifg.Tail.Filename, ",")

	m.Add(len(fileNames)) // 添加 go线程数量

	// 遍历取出对应的文件加上文件路径,启动go线程执行
	for _, fileName := range fileNames {
		file := []string{logpath, fileName}

		fmt.Printf("加载目录:%v\n", strings.Join(file, "/"))
		go redlogs(strings.Join(file, "/"), redisc)
	}

	m.Wait()


}
