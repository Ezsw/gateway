package lib

import (
	"bytes"
	"database/sql"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"strings"
)

var ConfBase *BaseConf
var ViperConfMap map[string]*viper.Viper // viper库用于处理应用程序的配置文件,使用ViperConfMap["config"] 获取名为 "config" 的配置文件
var DBMapPool map[string]*sql.DB
var GORMMapPool map[string]*gorm.DB
var DBDefaultPool *sql.DB

type BaseConf struct {
	DebugMode    string `mapstructure:"debug_mode"`
	TimeLocation string `mapstructure:"time_location"`
	Base         struct {
		DebugMode    string `mapstructure:"debug_mode"`
		TimeLocation string `mapstructure:"time_location"`
	} `mapstructure:"base"`
}

type MySQLConf struct {
	DriverName      string `mapstructure:"driver_name"`
	DataSourceName  string `mapstructure:"data_source_name"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time"`
}

type MysqlMapConf struct {
	List map[string]*MySQLConf `mapstructure:"list"`
}

// InitViperConf 初始化配置文件
func InitViperConf() error {
	f, err := os.Open(ConfEnvPath + "/") // 打开指定目录
	if err != nil {
		return err
	}
	fileList, err := f.Readdir(1024) //读取所有文件和子目录,限制最多读取1024个
	if err != nil {
		return err
	}
	for _, f0 := range fileList {
		if !f0.IsDir() {
			bts, err := ioutil.ReadFile(ConfEnvPath + "/" + f0.Name())
			if err != nil {
				return err
			}
			v := viper.New()
			v.SetConfigType("toml")
			v.ReadConfig(bytes.NewBuffer(bts))
			pathArr := strings.Split(f0.Name(), ".") // 将config.toml切分成["config", "toml"]
			if ViperConfMap == nil {
				ViperConfMap = make(map[string]*viper.Viper) // 创建一个新的 map 对象，并将其赋值给 ViperConfMap 变量。
			}
			ViperConfMap[pathArr[0]] = v
		}
	}
	return nil
}

// GetStringConf 获取get配置信息
func GetStringConf(key string) string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return ""
	}
	v, ok := ViperConfMap[keys[0]]
	if !ok {
		return ""
	}
	// strings.Join将 keys 数组中从下标 1 开始到最后一个元素（即 keys[1:len(keys)]）的所有元素使用 "." 进行连接，并返回连接后的字符串
	confString := v.GetString(strings.Join(keys[1:len(keys)], "."))
	return confString
}

// GetBoolConf 获取get配置信息
func GetBoolConf(key string) bool {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return false
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetBool(strings.Join(keys[1:len(keys)], "."))
	return conf
}

// GetIntConf 获取get配置信息
func GetIntConf(key string) int {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return 0
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetInt(strings.Join(keys[1:len(keys)], "."))
	return conf
}

// GetStringSliceConf 获取get配置信息
func GetStringSliceConf(key string) []string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetStringSlice(strings.Join(keys[1:len(keys)], "."))
	return conf
}

func InitBaseConf(path string) error {
	ConfBase = &BaseConf{}
	err := ParseConfig(path, ConfBase)
	if err != nil {
		return err
	}

	if ConfBase.DebugMode == "" {
		if ConfBase.Base.DebugMode != "" {
			ConfBase.DebugMode = ConfBase.Base.DebugMode
		} else {
			ConfBase.DebugMode = "debug"
		}
	}
	if ConfBase.TimeLocation == "" {
		if ConfBase.Base.TimeLocation != "" {
			ConfBase.TimeLocation = ConfBase.Base.TimeLocation
		} else {
			ConfBase.TimeLocation = "Asia/Shanghai"
		}
	}

	return nil
}
