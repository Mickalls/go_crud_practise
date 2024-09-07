package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"os"
)

var SqlConn *gorm.DB

const DRIVER = "mysql"

type conf struct {
	Url      string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"post"`
}

func InitMysql() (err error) {
	var mysqlConf conf
	mysqlConf.getConf()

	connetSqlUrl := BuildMySQLConnectionString(&mysqlConf)

	SqlConn, err = gorm.Open(DRIVER, connetSqlUrl)
	if err != nil {
		panic(err)
	}

	return SqlConn.DB().Ping()
}

func (c *conf) getConf() {
	//读取resources/application.yaml文件
	yamlFile, err := os.ReadFile("resources/application.yaml")
	// 若出现错误，打印错误提示并返回
	if err != nil {
		fmt.Println("Error reading YAML file:", err.Error())
		return
	}

	// 将读取的字符串转换成结构体conf
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println("Error unmarshaling YAML:", err.Error())
	}
}

/*
连接数据库需要的url:
username:password@protocol(address)/dbname?param=value
*/
func BuildMySQLConnectionString(cfg *conf) string {
	// 注意调整参数，确保安全性和编码（如密码中的特殊字符）
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.UserName, cfg.Password, cfg.Url, cfg.Port, cfg.DbName)
}

func Close() {
	SqlConn.Close()
}
