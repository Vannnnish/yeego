/**
 * Created by angelina on 2017/4/15.
 */

package yeesql_test

import (
	"fmt"
	"github.com/vannnnish/yeego"
	"github.com/vannnnish/yeego/yeesql"
	"strings"
	"testing"
)

func initTestDbTable() {
	yeesql.MustSetDbConfig(dbConf)
	yeesql.InitDbWithoutDbName()
	yeesql.MustCreateDb()
	yeesql.InitDb()
	yeesql.MustCreateTable(testTable)
}

func TestMustCreateDb(t *testing.T) {
	yeesql.MustSetDbConfig(dbConf)
	yeesql.InitDbWithoutDbName()
	yeesql.MustCreateDb()
}

func TestMustCreateTable(t *testing.T) {
	yeesql.MustSetDbConfig(dbConf)
	yeesql.InitDbWithoutDbName()
	yeesql.MustCreateDb()
	yeesql.InitDb()
	yeesql.MustCreateTable(testTable)
	yeego.OK(yeesql.MustIsTableExist(testTable.Name))
	yeesql.MustDropDb()
}

func TestMustSyncTable(t *testing.T) {
	initTestDbTable()
	testTable.FieldList = map[string]yeesql.DbType{
		"Id":       yeesql.DbTypeIntAutoIncrement,
		"Name":     yeesql.DbTypeString,
		"NewField": yeesql.DbTypeString,
	}
	yeesql.MustSyncTable(testTable)
	ret := yeesql.MustQueryOne("SHOW CREATE TABLE testTable")
	yeego.OK(strings.Contains(fmt.Sprint(ret), "Id"))
	yeego.OK(strings.Contains(fmt.Sprint(ret), "Name"))
	yeego.OK(strings.Contains(fmt.Sprint(ret), "Pwd"))
	yeego.OK(strings.Contains(fmt.Sprint(ret), "NewField"))
	yeego.OK(yeesql.MustIsTableExist(testTable.Name))
	yeesql.MustDropDb()
}

func TestMustForceSyncTable(t *testing.T) {
	initTestDbTable()
	testTable.FieldList = map[string]yeesql.DbType{
		"Id":       yeesql.DbTypeIntAutoIncrement,
		"Name":     yeesql.DbTypeString,
		"NewField": yeesql.DbTypeString,
	}
	yeesql.MustForceSyncTable(testTable)
	ret := yeesql.MustQueryOne("SHOW CREATE TABLE testTable")
	yeego.OK(strings.Contains(fmt.Sprint(ret), "Id"))
	yeego.OK(strings.Contains(fmt.Sprint(ret), "Name"))
	yeego.OK(!strings.Contains(fmt.Sprint(ret), "Pwd"))
	yeego.OK(strings.Contains(fmt.Sprint(ret), "NewField"))
	yeego.OK(yeesql.MustIsTableExist(testTable.Name))
	yeesql.MustDropDb()
}
