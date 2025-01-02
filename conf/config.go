package conf

import (
	"demoProject4mall/dao"
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppMode  string
	HttpPort string
	DB       string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RedisDb     string
	RedisAddr   string
	redisPw     string
	RedisDbName string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	Host        string
	ProductPath string
	AvatarPath  string
)

func Init() {
	//本地读取环境变量   go get gopkg.in/ini.v1
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	LoadMysql(file)
	LoadRedis(file)
	LoadEmail(file)
	LoadPath(file)
	//mysql读（8）主
	pathRead := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True"}, "")
	//mysql写（2）从 主从复制
	pathWrite := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True"}, "")
	fmt.Println("pathRead:", pathRead)
	fmt.Println("pathWrite:", pathWrite)
	dao.Database(pathRead, pathWrite)
}
func LoadServer(file *ini.File) {
	//读取server下的AppMode转换为字符串
	AppMode = file.Section("server").Key("AppMode").String()
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8080")

}
func LoadMysql(file *ini.File) {
	DB = file.Section("mysql").Key("DB").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").MustString("3306")
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()

}
func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	redisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()

}
func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()

}
func LoadPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()

}
