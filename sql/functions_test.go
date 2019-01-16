/**
 * Created by angelina on 2017/4/15.
 */

package sql_test

import (
	"github.com/vannnnish/yeego"
	"github.com/vannnnish/yeego/sql"
	"testing"
)

var (
	dbConf = &sql.DbConfig{
		UserName: "root",
		Password: "root",
		Host:     "127.0.0.1",
		Port:     "3306",
		DbName:   "yeeSql_test",
	}
	testTable = sql.Table{
		Name: "testTable",
		FieldList: map[string]sql.DbType{
			"Id":   sql.DbTypeIntAutoIncrement,
			"Name": sql.DbTypeString,
			"Pwd":  sql.DbTypeString,
		},
		PrimaryKey: "Id",
		UniqueKey: [][]string{
			[]string{"Id"},
		},
		NotNull: []string{"Name", "Pwd"},
	}
	tomlData = `
				[[testTable]]
				Id = "1"
				Name = "👮👮👮"
				Pwd = "111"
				[[testTable]]
				Id = "2"
				Name = "angelina2"
				Pwd = "222"
				[[testTable]]
				Id = "3"
				Name = "angelina3"
				Pwd = "333"
			`
)

func setTestTableData() {
	sql.MustSetTableDataToml(tomlData)
}

func TestQuery(t *testing.T) {
	initTestDbTable()
	setTestTableData()
	data, err := sql.Query("SELECT * FROM testTable")
	yeego.Equal(err, nil)
	yeego.Equal(len(data), 3)
	yeego.Equal(data[0]["Id"], "1")
	setTestTableData()
	infoA, _ := sql.Query("SELECT * FROM testTable LIMIT 1")
	infoB, _ := sql.QueryOne("SELECT * FROM testTable")
	yeego.Equal(len(infoA), 1)
	yeego.Equal(infoA[0]["Id"], infoB["Id"])
	yeego.Equal(infoA[0]["Name"], infoB["Name"])
}
func TestInsert(t *testing.T) {
	initTestDbTable()
	setTestTableData()
	id, err := sql.Insert("testTable", map[string]string{
		"Id":   "4",
		"Name": "👮",
		"Pwd":  "444",
	})
	yeego.Equal(id, 4)
	yeego.Equal(err, nil)
	info, err := sql.QueryOne("SELECT * FROM testTable WHERE Id = 4")
	yeego.Equal(err, nil)
	yeego.Equal(info["Name"], "👮")
	yeego.Equal(info["Pwd"], "444")
}

func TestUpdateByID(t *testing.T) {
	initTestDbTable()
	setTestTableData()
	err := sql.UpdateByID("testTable", "Id", map[string]string{
		"Id":   "1",
		"Name": "changed",
		"Pwd":  "changed",
	})
	yeego.Equal(err, nil)
	info, err := sql.QueryOne("SELECT * FROM testTable WHERE Id = 1")
	yeego.Equal(err, nil)
	yeego.Equal(info["Name"], "changed")
	yeego.Equal(info["Pwd"], "changed")
}

func TestDeleteByID(t *testing.T) {
	initTestDbTable()
	setTestTableData()
	err := sql.DeleteByID("testTable", "Id", "1")
	yeego.Equal(err, nil)
	info, err := sql.GetOneWhere("testTable", "Id", "1")
	yeego.NotEqual(err, nil)
	yeego.Equal(info, nil)
}

func TestGetOneWhere(t *testing.T) {
	initTestDbTable()
	setTestTableData()
	info, err := sql.GetOneWhere("testTable", "Id", "1")
	yeego.Equal(err, nil)
	yeego.Equal(info["Name"], "angelina1")
}

func TestGetAllInTable(t *testing.T) {
	initTestDbTable()
	setTestTableData()
	all, err := sql.GetAllInTable("testTable")
	yeego.Equal(err, nil)
	yeego.Equal(len(all), 3)
}

func TestIsExist(t *testing.T) {
	initTestDbTable()
	setTestTableData()
	yeego.Equal(sql.IsExist("testTable", map[string]string{
		"Id":   "1",
		"Name": "angelina1",
	}), true)
	yeego.Equal(sql.IsExist("testTable", map[string]string{
		"Id":   "2",
		"Name": "angelina1",
	}), false)
}
