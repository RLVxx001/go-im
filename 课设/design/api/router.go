package api

import (
	"design/config"
	"design/utils/database_handler"
	"fmt"
	"log"
)

// Databases 结构体
type Databases struct {
}

// 配置文件全局对象
var AppConfig = &config.Configuration{}

// 根据配置文件创建数据库
func CreateDBs() *Databases {
	cfgFile := "./config/config.yaml"
	conf, err := config.GetAllConfigValues(cfgFile)
	AppConfig = conf
	if err != nil {
		log.Fatalf("读取配置文件失败. %v", err.Error())
	}
	m := AppConfig.DatabaseSettings
	var dns = fmt.Sprintf("%s:%s@%s/%s?%s", m.Username, m.Password, m.DatabaseURIL, m.DatabaseName, m.DatabaseURIR)

	db := database_handler.NewMySQLDB(dns)
	fmt.Printf("%v", db)
	return &Databases{}
}
