package lib

import (
	"fmt"
	"log"
	"time"
)

var TimeLocation *time.Location
var TimeFormat = "2006-01-02 15:04:05" // go格式化时间123456便于记忆

// InitModule 标准流程
// 函数传入配置文件 InitModule("./conf/dev/")
func InitModule(configPath string) error {
	return initModule(configPath, []string{"base", "mysql", "redis"})
}

// initModule 加载配置累
func initModule(configPath string, modules []string) error {
	log.Println("------------------------------------------------------------------------")
	log.Printf("[INFO]  config=%s\n", configPath)
	log.Printf("[INFO] %s\n", " start loading resources.")

	// 解析配置文件目录
	if err := ParseConfPath(configPath); err != nil {
		return err
	}

	//初始化配置文件
	if err := InitViperConf(); err != nil {
		return err
	}

	// 加载base配置
	if InArrayString("base", modules) {
		if err := InitBaseConf(GetConfPath("base")); err != nil {
			fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitBaseConf:"+err.Error())
		}
	}

	// 加载mysql配置并初始化实例
	if InArrayString("mysql", modules) {
		if err := InitDBPool(GetConfPath("mysql_map")); err != nil {
			fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitDBPool:"+err.Error())
		}
	}

	// 设置时区
	if location, err := time.LoadLocation(ConfBase.TimeLocation); err != nil {
		return err
	} else {
		TimeLocation = location
	}

	log.Printf("[INFO] %s\n", " success loading resources.")
	log.Println("------------------------------------------------------------------------")
	return nil
}

// Destroy 公共销毁函数
func Destroy() {
	log.Println("------------------------------------------------------------------------")
	log.Printf("[INFO] %s\n", " start destroy resources.")
	CloseDB()

	log.Printf("[INFO] %s\n", " success destroy resources.")
}
