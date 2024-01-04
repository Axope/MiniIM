package pool

import (
	"MiniIM/configs"
	"MiniIM/internal/models"
	"database/sql"
	"fmt"

	"MiniIM/pkg/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func autoMigrate() {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Logger.Error("sql AutoMigrate User error")
	}
	db.AutoMigrate(&models.Friend{})
	if err != nil {
		log.Logger.Error("sql AutoMigrate Friend error")
	}
	db.AutoMigrate(&models.Group{})
	if err != nil {
		log.Logger.Error("sql AutoMigrate Group error")
	}
	db.AutoMigrate(&models.GroupMember{})
	if err != nil {
		log.Logger.Error("sql AutoMigrate GroupMember error")
	}
}

func Init() {
	sqlConfig := configs.GetConfig().Mysql

	username := sqlConfig.Username
	pwd := sqlConfig.Password
	host := sqlConfig.Host
	port := sqlConfig.Port
	dbname := sqlConfig.DBname
	timeout := sqlConfig.Timeout

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		username, pwd, host, port, dbname, timeout)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Logger.Panic(err.Error())
	}
	log.Logger.Info("mysql init success")

	// 设置连接池
	var sqlDB *sql.DB
	sqlDB, err = db.DB()
	if err != nil {
		log.Logger.Warn(err.Error())
		return
	}
	sqlDB.SetMaxIdleConns(10)  // 连接池最大空闲连接数
	sqlDB.SetMaxOpenConns(100) // 连接池最大可打开连接数

	log.Logger.Info("mysql pool init success")
	autoMigrate()
}

func GetDB() *gorm.DB {
	return db
}
