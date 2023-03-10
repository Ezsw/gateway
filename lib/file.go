package lib

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

// 与解析文件相关的方法

var ConfEnvPath string //配置文件夹
var ConfEnv string     //配置环境名 比如：dev prod test

// ParseConfPath 解析配置文件目录   如：config=conf/dev/base.json 	ConfEnvPath=conf/dev	ConfEnv=dev
func ParseConfPath(config string) error {
	path := strings.Split(config, "/")              //将字符串分割成切片
	prefix := strings.Join(path[:len(path)-1], "/") // 将除最后一个元素外的切片合成一个字符串
	ConfEnvPath = prefix
	ConfEnv = path[len(path)-2]
	return nil
}

// GetConfEnv 获取配置环境名
func GetConfEnv() string {
	return ConfEnv
}

func GetConfPath(fileName string) string {
	return ConfEnvPath + "/" + fileName + ".toml"
}

func GetConfFilePath(fileName string) string {
	return ConfEnvPath + "/" + fileName
}

func ParseConfig(path string, conf interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Open config %v fail, %v", path, err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Read config fail, %v", err)
	}

	v := viper.New()
	v.SetConfigType("toml")
	v.ReadConfig(bytes.NewBuffer(data))
	if err := v.Unmarshal(conf); err != nil {
		return fmt.Errorf("Parse config fail, config:%v, err:%v", string(data), err)
	}
	return nil
}
