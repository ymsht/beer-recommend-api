package db

import (
	"beer-recommend-api/model"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/gorp.v1"
)

// MysqlConfig DB設定
type MysqlConfig struct {
	Host     string `envconfig:"DB_HOST"`
	Port     int    `envconfig:"DB_PORT"`
	User     string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
	Name     string `envconfig:"DB_NAME"`
}

// Init 初期化します
func Init() *gorp.DbMap {
	dbmap := getDbMap()
	return dbmap
}

// getDbMap DbMapを取得します
func getDbMap() *gorp.DbMap {
	dsn, err := getDsn()
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	dbmap.AddTableWithName(model.Review{}, "review").SetKeys(true, "ReviewId")
	dbmap.AddTableWithName(model.Flavor{}, "flavor")
	dbmap.AddTableWithName(model.Country{}, "country")
	dbmap.AddTableWithName(model.Area{}, "area")
	dbmap.AddTableWithName(model.User{}, "user")

	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	return dbmap
}

func getDsn() (string, error) {
	var conf MysqlConfig
	err := envconfig.Process("", &conf)
	if err != nil {
		return "", nil
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Name)

	return dsn, nil
}
