/**
 * Created by angelina on 2017/4/15.
 */

package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var dbWithoutDbName *sqlx.DB
var dbConfig DbConfig

// InitDb
// 初始化DB
func InitDb() {
	MustVerifyDbConfig()
	if db == nil {
		db = sqlx.MustConnect("mysql", dbConfig.GetDsn())
	}
	if dbWithoutDbName == nil {
		dbWithoutDbName = sqlx.MustConnect("mysql", dbConfig.GetDsnWithoutDbName())
	}
	db.SetMaxIdleConns(0)
}

func InitDbWithoutDbName() {
	MustVerifyDbConfig()
	if dbWithoutDbName == nil {
		dbWithoutDbName = sqlx.MustConnect("mysql", dbConfig.GetDsnWithoutDbName())
	}
}

func GetDb() *sqlx.DB {
	if db == nil {
		panic("请先 InitDb")
	}
	return db
}

func GetDbWithoutDbName() *sqlx.DB {
	if dbWithoutDbName == nil {
		panic("请先 InitDb")
	}
	return dbWithoutDbName
}
