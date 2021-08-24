package models

import (
	"fmt"
	"github.com/13808796047/go-gin-example/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"parmary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("fail to get section 'database':%v", err)
	}
	dbType = sec.Key("TYPE").String()
	password = sec.Key("PASSWORD").String()
	user = sec.Key("USER").String()
	dbName = sec.Key("NAME").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	// 连接数据库
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err,user,password,host,dbName)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	//CloseDB()
}
func CloseDB() {
	defer db.Close()
}
