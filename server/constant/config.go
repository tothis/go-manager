package constant

import (
	"fmt"
	"github.com/spf13/viper"
	"go-manager/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	appConfig = initConfig()
	Db        = initMySQLConfig()
)

// 读取配置文件config
type Config struct {
	MySQL MySQLConfig
}

type MySQLConfig struct {
	Host     string
	Port     uint16
	DbName   string
	Arg      string
	UserName string
	Password string
}

// 初始化配置
func initConfig() Config {
	// 把配置文件读取到结构体上
	var config Config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	// 将配置文件绑定到config上
	viper.Unmarshal(&config)
	fmt.Println("config : ", config)
	return config
}

func initMySQLConfig() *gorm.DB {
	m := appConfig.MySQL
	link := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.UserName, m.Password, m.Host, m.Port, m.DbName, m.Arg)
	config := &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			// 创建表时候去掉's'后缀
			SingularTable: true,
		},
	}
	db, err := gorm.Open(mysql.Open(link), config)
	if err != nil {
		panic("连接数据库失败")
	}
	initMySQLTable(db)
	return db
}

// 自动建表
func initMySQLTable(db *gorm.DB) {
	db.AutoMigrate(model.User{})
}
