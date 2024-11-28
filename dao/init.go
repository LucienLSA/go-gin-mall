package dao

import (
	"ginmall/conf"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func InitMySQL() {
	// 初始化GORM日志配置
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // Slow SQL threshold
	// 		LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
	// 		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	// 		Colorful:                  false,       // Disable color
	// 	},
	// )
	mConfig := conf.Config.MySql["default"]
	// 读写分离
	pathRead := strings.Join([]string{mConfig.UserName, ":", mConfig.Password, "@tcp(", mConfig.DbHost, ":", mConfig.DbPort, ")/", mConfig.DbName, "?charset=" + mConfig.Charset + "&parseTime=true"}, "")
	pathWrite := strings.Join([]string{mConfig.UserName, ":", mConfig.Password, "@tcp(", mConfig.DbHost, ":", mConfig.DbPort, ")/", mConfig.DbName, "?charset=" + mConfig.Charset + "&parseTime=true"}, "")

	var ormLogger logger.Interface
	if conf.Config.System.AppEnv == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       pathRead,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	// Error
	// if connString == "" || err != nil {
	// 	logging.LogrusObj.Error("mysql lost: %v", err)
	// }
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	logging.LogrusObj.Error("mysql lost: %v", err)
	// 	panic(err)
	// }
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db
	_ = DB.Use(dbresolver.
		Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(pathRead)},
			Replicas: []gorm.Dialector{mysql.Open(pathWrite)},
			Policy:   dbresolver.RandomPolicy{},
		}))
	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(10)
	//打开
	sqlDB.SetMaxOpenConns(50)
	DB = db

	migration()
}
