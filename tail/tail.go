package tail

// go get github.com/hpcloud/tail

import (
	"fmt"
	// "os"

	"github.com/hpcloud/tail"
)

// tail读取文件
func TailRead(fileName string, getoffset int64) *tail.Tail {
	tailcConfig := tail.Config{
		ReOpen:    true,  // 是否打开
		Follow:    true,  // 是否跟随
		MustExist: false, // 文件不存在不报错
		Poll:      true,
		/* 开始读入文件
		 */
		Location: &tail.SeekInfo{Offset: getoffset, Whence: 1},
	}

	tail, err := tail.TailFile(fileName, tailcConfig)
	if err != nil {
		fmt.Printf("tail读取文件失败:%v\n", err)
		return tail
	}

	return tail
}
