package conf

import (
	"fmt"
	"strconv"

	"gopkg.in/ini.v1"
)

type WebConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

type MySqlConfig struct {
	Ip       string `ini:"ip"`
	Port     int    `ini:"port"`
	UserName string `ini:"username"`
	PassWord string `ini:"password"`
	DBName   string `ini:"dbname"`
	ConPool  int    `ini:"conpool"`
}

func (mqc *MySqlConfig) GetMySqlConInfo() string {
	return mqc.UserName + ":" + mqc.PassWord + "@tcp(" + mqc.Ip + ":" + strconv.Itoa(mqc.Port) + ")/" + mqc.DBName
}

type RedisConfig struct {
	Ip       string `ini:"ip"`
	Port     int    `ini:"port"`
	PassWord string `ini:"password"`
	ConPool  int    `ini:"conpool"`
}

type LogConfig struct {
	FileName       string `ini:"fileName"`
	MaxFileSize    int    `ini:"maxFileSize"`
	MaxBackupIndex int    `ini:"maxBackupIndex"`
	MaxAge         int    `ini:"maxAge"`
	Compress       bool   `ini:"compress"`
	Level          int    `ini:"level"`
}

var (
	WebCfg   WebConfig
	MySqlCfg MySqlConfig
	RedisCfg RedisConfig
	LogCfg   LogConfig
)

func ReadConf(confFile string) error {
	cfg, err := ini.Load(confFile)
	if err != nil {
		return fmt.Errorf("load ini failed: %s", err.Error())
	}
	err = cfg.Section("web").MapTo(&WebCfg)
	if err != nil {
		return fmt.Errorf("read web ini failed: %s", err.Error())
	}
	err = cfg.Section("mysql").MapTo(&MySqlCfg)
	if err != nil {
		return fmt.Errorf("read web mysql failed: %s", err.Error())
	}
	err = cfg.Section("redis").MapTo(&RedisCfg)
	if err != nil {
		return fmt.Errorf("read web redis failed: %s", err.Error())
	}
	err = cfg.Section("log").MapTo(&LogCfg)
	if err != nil {
		return fmt.Errorf("read log ini failed: %s", err.Error())
	}
	fmt.Println(MySqlCfg.GetMySqlConInfo())
	fmt.Println(WebCfg)
	fmt.Println(MySqlCfg)
	fmt.Println(RedisCfg)
	fmt.Println(LogCfg)
	return nil
}
