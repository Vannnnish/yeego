/**
 * Created by angelina on 2017/4/16.
 */

package yeeSql

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jmoiron/sqlx"
	"strings"
)

// MustSetTableDataToml
// 为数据库table设置数据 toml格式
// 需要注意：
// 1.会清除表的数据
// 2.需要设置好自增字段的值（比数据库最大的+1）
func MustSetTableDataToml(data string) {
	tableData := make(map[string][]map[string]string)
	_, err := toml.Decode(data, &tableData)
	if err != nil {
		panic(err)
	}
	if len(tableData) == 0 {
		panic("[MustSetTableDataToml] want to set data but data is empty?")
	}
	if err := setTableDataToml(tableData); err != nil {
		panic(err)
	}
}

func setTableDataToml(data map[string][]map[string]string) error {
	db := GetDb()
	tx := db.MustBegin()
	err := setTablesDataTransaction(data, tx)
	if err != nil {
		errRoll := tx.Rollback()
		if errRoll != nil {
			return fmt.Errorf("[setTableDataToml] error [transaction] %s,[rollback] %s", err, errRoll)
		}
		return err
	}
	return tx.Commit()
}

func setTablesDataTransaction(data map[string][]map[string]string, tx *sqlx.Tx) error {
	for tableName, tableData := range data {
		truncateSql := fmt.Sprintf("TRUNCATE `%s`", tableName)
		if _, err := tx.Exec(truncateSql); err != nil {
			return err
		}
		for _, row := range tableData {
			colNameList := []string{}
			placeHolderNum := len(row)
			valueList := []interface{}{}
			for name, value := range row {
				colNameList = append(colNameList, name)
				valueList = append(valueList, value)
			}
			sqlColNamePart := "`" + strings.Join(colNameList, "`, `") + "`"
			sqlValuePart := strings.Repeat("?, ", placeHolderNum-1) + "?"
			insertSql := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)", tableName, sqlColNamePart, sqlValuePart)
			if _, err := tx.Exec(insertSql, valueList...); err != nil {
				return err
			}
		}
	}
	return nil
}
