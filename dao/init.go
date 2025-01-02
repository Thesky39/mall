package dao

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

var _db *gorm.DB

func Database(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,  //string类型字段默认长度
		DisableDatetimePrecision:  true, //禁止datetime精度
		DontSupportRenameIndex:    true, //重命名索引，就要把索引删除再重建
		DontSupportRenameColumn:   true, //用change重命名列
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		//gorm的配置 配置日志器
		Logger: ormLogger,
		//配置命名策略
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	//获取db对象的原生数据库连接池（即可以使用sql）
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)  //设置连接池
	sqlDB.SetMaxOpenConns(100) //打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour)
	_db = db

	//主从配置  go get gorm.io/plugin/dbresolver
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(connWrite)},                       //写
		Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connWrite)}, //读
	}))
	migration()
}
func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db
}
