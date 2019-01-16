/**
 * Created by angelina on 2017/4/15.
 */

package sql_test

import (
	"fmt"
	"github.com/vannnnish/yeego"
	"github.com/vannnnish/yeego/sql"
	"strings"
	"testing"
)

func initTestDbTable() {
	sql.MustSetDbConfig(dbConf)
	sql.InitDbWithoutDbName()
	sql.MustCreateDb()
	sql.InitDb()
	sql.MustCreateTable(testTable)
}

func TestMustCreateDb(t *testing.T) {
	sql.MustSetDbConfig(dbConf)
	sql.InitDbWithoutDbName()
	sql.MustCreateDb()
}

func TestMustCreateTable(t *testing.T) {
	sql.MustSetDbConfig(dbConf)
	sql.InitDbWithoutDbName()
	sql.MustCreateDb()
	sql.InitDb()
	sql.MustCreateTable(testTable)
	yeego.OK(sql.MustIsTableExist(testTable.Name))
	sql.MustDropDb()
}

func TestMustSyncTable(t *testing.T) {
	initTestDbTable()
	testTable.FieldList = map[string]sql.DbType{
		"Id":       sql.DbTypeIntAutoIncrement,
		"Name":     sql.DbTypeString,
		"NewField": sql.DbTypeString,
	}
	sql.MustSyncTable(testTable)
	ret := sql.MustQueryOne("SHOW CREATE TABLE testTable")
	yeego.OK(strings.Contains(fmt.Sprint(ret), "Id"))
	yeego.OK(strings.Contains(fmt.Sprint(ret), "Name"))
	yeego.OK(strings.Contains(fmt.Sprint(ret), "Pwd"))
	yeego.OK(strings.Contains(fmt.Sprint(ret), "NewField"))
	yeego.OK(sql.MustIsTableExist(testTable.Name))
	sql.MustDropDb()
}

func TestMustForceSyncTable(t *testing.T) {
	initTestDbTable()
	testTable.FieldList = map[string]sql.DbType{
		"Id":       sql.DbTypeIntAutoIncrement,
		"Name":     sql.DbTypeString,
		"NewField": sql.DbTypeString,
	}
	sql.MustForceSyncTable(testTable)
	ret := sql.MustQueryOne("SHOW CREATE TABLE testTable")
	yeego.OK(strings.Contains(fmt.Sprint(ret), "Id"))
	yeego.OK(strings.Contains(fmt.Sprint(ret), "Name"))
	yeego.OK(!strings.Contains(fmt.Sprint(ret), "Pwd"))
	yeego.OK(strings.Contains(fmt.Sprint(ret), "NewField"))
	yeego.OK(sql.MustIsTableExist(testTable.Name))
	sql.MustDropDb()
}
