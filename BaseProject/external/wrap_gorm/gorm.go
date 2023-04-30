package wrap_gorm

import (
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
	"log"
	"os"
	"time"

	MySQLDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db     *gorm.DB
	config *MySQLDriver.Config
	vp     = wrap_viper.GetViper()
)

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	err := ResetConfDB()
	if err != nil {
		panic(err)
	}
	return db
}

func ResetConfDB() error {
	if config == nil {
		SetDbConfig(vp.GetString("DB_HOST"), vp.GetInt("DB_PORT"), vp.GetString("DB_NAME"), vp.GetString("DB_USER"), vp.GetString("DB_PASSWORD"), vp.GetString("DB_LOCATION"))
	}
	err := connectDB()
	if err != nil {
		panic(err)
	}
	return nil
}

func SetDbConfig(DbHost string, DbPort int, DbName string, DbUser string, DbPassword string, DbLocation string) {
	locat, err := time.LoadLocation(DbLocation)
	if err != nil {
		panic(err)
	}
	config = MySQLDriver.NewConfig()

	config.User = DbUser
	config.Passwd = DbPassword
	config.Net = "tcp"
	config.Addr = fmt.Sprintf("%s:%v", DbHost, DbPort)
	config.DBName = DbName
	config.Loc = locat
	config.ParseTime = true
	config.Params = map[string]string{"charset": "utf8mb4"}
	config.AllowNativePasswords = true
}

func connectDB() error {
	var err error

	dbUrl := config.FormatDSN()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
		},
	)

	db, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{Logger: newLogger})
	if err != nil {
		db = nil
		config = nil
		panic(err)
	}

	sql, err := db.DB()
	if err != nil {
		db = nil
		config = nil
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sql.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sql.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sql.SetConnMaxLifetime(time.Hour)

	return nil
}
