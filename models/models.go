package models

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/itachilee/gin/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DbSetting struct {
	Host     string
	Port     int
	Username string
	Password string
	Schema   string
}

// Setup initializes the database instance
func Setup() error {
	var err error

	if global.DatabaseSetting == nil {
		return errors.New("数据库配置不能为空")
	}
	switch strings.ToLower(global.DatabaseSetting.DBType) {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
			global.DatabaseSetting.UserName,
			global.DatabaseSetting.Password,
			global.DatabaseSetting.Host,
			global.DatabaseSetting.DBName,
			global.DatabaseSetting.Charset,
			global.DatabaseSetting.ParseTime,
			global.DatabaseSetting.Loc,
		)
		global.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{

			Logger: logger.Default.LogMode(logger.Info),
			// Logger: logger.Default.LogMode(logger.Silent),
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   global.DatabaseSetting.TablePrefix, // table name prefix, table for `User` would be `t_users`
				SingularTable: false,                              // use singular table name, table for `User` would be `user` with this option enabled
				NoLowerCase:   true,                               // skip the snake_casing of names
				// NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
			},
		})
	default:
		return errors.New("未实现该类型的数据库链接")
	}

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	return err
	// db.AutoMigrate(&Address{}, &GoodsDetail{})
}
