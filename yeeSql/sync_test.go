/**
 * Created by angelina on 2017/4/15.
 */

package yeeSql_test

import (
	"fmt"
	"yeego"
	"yeego/yeeSql"
	"strings"
	"testing"
)

func initTestDbTable() {
	yeeSql.MustSetDbConfig(dbConf)
	yeeSql.InitDbWithoutDbName()
	yeeSql.MustCreateDb()
	yeeSql.InitDb()
	yeeSql.MustCreateTable(testTable)
}

func TestMustCreateDb(t *testing.T) {
	yeeSql.MustSetDbConfig(dbConf)
	yeeSql.InitDbWithoutDbName()
	yeeSql.MustCreateDb()
}

func TestMustCreateTable(t *testing.T) {
	yeeSql.MustSetDbConfig(dbConf)
	yeeSql.InitDbWithoutDbName()
	yeeSql.MustCreateDb()
	yeeSql.InitDb()
	yeeSql.MustCreateTable(testTable)
	yeego.OK(yeeSql.MustIsTableExist(testTable.Name))
	yeeSql.MustDropDb()
}

func TestMustSyncTable(t *testing.T) {
	initTestDbTable()
	testTable.FieldList = map[string]yeeSql.DbType{
		"Id":       yeeSql.DbTypeIntAutoIncrement,
		"Name":     yeeSql.DbTypeString,
		"NewField": yeeSql.DbTypeString,
	}
	yeeSql.MustSyncTable(testTable)
	ret := yeeSql.MustQueryOne("SHOW CREATE TABLE testTable")
	yeego.OK(strings.Contains(fmt.Sprint(ret), "Id"))
	yeego.OK(strings.Contains(fmt.Sprint(ret), "Name"))
	yeego.OK(strings.Contains(fmt.Sprint(ret), "Pwd"))
	yeego.OK(strings.Contains(fmt.Sprint(ret), "NewField"))
	yeego.OK(yeeSql.MustIsTableExist(testTable.Name))
	yeeSql.MustDropDb()
}

func TestMustForceSyncTable(t *testing.T) {
	initTestDbTable()
	testTable.FieldList = map[string]yeeSql.DbType{
		"Id":       yeeSql.DbTypeIntAutoIncrement,
		"Name":     yeeSql.DbTypeString,
		"NewField": yeeSql.DbTypeString,
	}
	yeeSql.MustForceSyncTable(testTable)
	ret := yeeSql.MustQueryOne("SHOW CREATE TABLE testTable")
	yeego.OK(strings.Contains(fmt.Sprint(ret), "Id"))
	yeego.OK(strings.Contains(fmt.Sprint(ret), "Name"))
	yeego.OK(!strings.Contains(fmt.Sprint(ret), "Pwd"))
	yeego.OK(strings.Contains(fmt.Sprint(ret), "NewField"))
	yeego.OK(yeeSql.MustIsTableExist(testTable.Name))
	yeeSql.MustDropDb()
}
