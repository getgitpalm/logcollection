package conf

// go get gopkg.in/ini.v1

import (
	"errors"
	"fmt"

	"gopkg.in/ini.v1"
)

type configArgs struct {
	Redis redis `ini:"redis"`
	Tail  tail  `ini:"tail"`
}

type redis struct {
	Address string `ini:"address"`
	Port    string `ini:"port"`
	Passwd  string `ini:"passwd"`
}

type tail struct {
	Logpath  string `ini:"logpath"`
	Filename string `ini:"filename"`
}

// 解析ini文件

func Reloadconf() (configArgs, error) {

	configArgsObj := configArgs{}

	cf, err := ini.Load("conf/config.ini")
	if err != nil {
		fmt.Printf("加载config.ini失败:%v\n", err)
		errors.New("加载config.ini失败")
		return configArgsObj, err
	}

	err = cf.MapTo(&configArgsObj)
	if err != nil {
		fmt.Printf("解析config.ini失败:%v\n", err)
		errors.New("解析config.ini失败")
		return configArgsObj, err
	}

	return configArgsObj, err

}
