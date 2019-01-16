/**
 * Created by angelina on 2017/4/15.
 */

package sql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/vannnnish/yeego/sql/MysqlAst"
	"strings"
)

// RowsToMapSlice
// 将*sql.Row读取到[]map[string]string
func RowsToMapSlice(rows *sql.Rows) (output []map[string]string) {
	if rows == nil {
		return
	}
	// 获取table的字段数量
	columns, err := rows.Columns()
	if err != nil {
		return
	}
	lenColumn := len(columns)
	for rows.Next() {
		rowArray := make([]interface{}, lenColumn)
		// 用*RawByte包装每一个值
		for k1 := range rowArray {
			var s sql.RawBytes
			rowArray[k1] = &s
		}
		// 读取这个row的全部值
		if err := rows.Scan(rowArray...); err != nil {
			return
		}
		rowMap := make(map[string]string)
		// 解包装，将全部字段取出来
		for rowIndex, rowName := range columns {
			rowMap[rowName] = string(*(rowArray[rowIndex].(*sql.RawBytes)))
		}
		// 这个row扔到output中
		output = append(output, rowMap)
	}
	return
}

func RowsToMapSliceFirst(rows *sql.Rows) (map[string]string, error) {
	out := RowsToMapSlice(rows)
	if len(out) > 0 {
		return out[0], nil
	}
	return nil, errors.New("not found row")
}

// Query
// 查询语句
func Query(query string, args ...interface{}) (output []map[string]string, err error) {
	rows, err := GetDb().Query(GetDb().Rebind(query), args...)
	colorPrint("[sql.Query]:Query=["+GetDb().Rebind(query)+"] args=", args)
	if err != nil {
		return nil, fmt.Errorf("[Query] sql: [%s] data: [%s] err:[%s]",
			query, strings.Join(argsInterfaceToString(args), ","), err.Error())
	}
	defer rows.Close()
	// 获取table的字段数量
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	lenColumn := len(columns)
	for rows.Next() {
		rowArray := make([]interface{}, lenColumn)
		// 用*RawByte包装每一个值
		for k1 := range rowArray {
			var s sql.RawBytes
			rowArray[k1] = &s
		}
		// 读取这个row的全部值
		if err := rows.Scan(rowArray...); err != nil {
			return nil, err
		}
		rowMap := make(map[string]string)
		// 解包装，将全部字段取出来
		for rowIndex, rowName := range columns {
			rowMap[rowName] = string(*(rowArray[rowIndex].(*sql.RawBytes)))
		}
		// 这个row扔到output中
		output = append(output, rowMap)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return
}

func MustQuery(query string, args ...interface{}) (output []map[string]string) {
	output, err := Query(query, args...)
	if err != nil {
		panic(err)
	}
	return
}

// QueryOne
// 查询一条数据
// 如果有多条，则返回第一条
// 找不到返回err
func QueryOne(query string, args ...interface{}) (output map[string]string, err error) {
	list, err := Query(query, args...)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("no row find")
	}
	output = list[0]
	return output, err
}

func MustQueryOne(query string, args ...interface{}) map[string]string {
	out, err := QueryOne(query, args...)
	if err != nil {
		panic(err)
	}
	return out
}

// Exec
// 执行语句
func Exec(query string, args ...interface{}) (sql.Result, error) {
	colorPrint("[sql.Exec]:Query=["+GetDb().Rebind(query)+"] args=", args)
	return GetDb().Exec(GetDb().Rebind(query), args...)
}

func MustExec(query string, args ...interface{}) {
	_, err := Exec(query, args...)
	if err != nil {
		panic(err)
	}
}

// Insert
// 插入语句
// 通常返回的是主键的自增id
func Insert(tableName string, row map[string]string) (lastInsertID int, err error) {
	keyList := []string{}
	valueList := []string{}
	for key, value := range row {
		keyList = append(keyList, key)
		valueList = append(valueList, value)
	}
	keyStr := "`" + strings.Join(keyList, "`,`") + "`"
	valueStr := strings.Repeat("?,", len(row)-1) + "?"
	sql := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)", tableName, keyStr, valueStr)
	result, err := Exec(sql, argsStringToInterface(valueList...)...)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	lastInsertID = int(id)
	return
}

func MustInsert(tableName string, row map[string]string) (lastInsertID int) {
	lastInsertID, err := Insert(tableName, row)
	if err != nil {
		panic(err)
	}
	return
}

// UpdateByID
// 通过primaryKeyName的值去更新数据
// 通常是主键id
// 假设通过id=5去更新，id=5这条数据不存在，返回nil，不会报错
func UpdateByID(tableName string, primaryKeyName string, row map[string]string) error {
	keyList := []string{}
	valueList := []string{}
	var primaryValue string
	for key, value := range row {
		if primaryKeyName == key {
			primaryValue = value
			continue
		}
		keyList = append(keyList, "`"+key+"`=?")
		valueList = append(valueList, value)
	}
	if primaryValue == "" {
		return fmt.Errorf("primaryKey %s not set", primaryKeyName)
	}
	valueList = append(valueList, primaryValue)
	updateStr := strings.Join(keyList, ",")
	// UPDATE User SET Name=?,Pwd=? WHERE id = ?
	sql := fmt.Sprintf("UPDATE `%s` SET %s WHERE `%s` = ?", tableName, updateStr, primaryKeyName)
	_, err := Exec(sql, argsStringToInterface(valueList...)...)
	if err != nil {
		return err
	}
	return nil
}

func MustUpdateByID(tableName string, primaryKeyName string, row map[string]string) {
	err := UpdateByID(tableName, primaryKeyName, row)
	if err != nil {
		panic(err)
	}
}

// DeleteByID
// 通过字段删除某条值
func DeleteByID(tableName, fieldName, value string) error {
	deleteSql := fmt.Sprintf("DELETE FROM `%s` WHERE `%s` = ?", tableName, fieldName)
	_, err := Exec(deleteSql, value)
	return err
}

func MustDeleteByID(tableName, fieldName, value string) {
	err := DeleteByID(tableName, fieldName, value)
	if err != nil {
		panic(err)
	}
}

// GetOneWhere
// 通过某个字段的值查找一条数据
func GetOneWhere(tableName, fieldName, value string) (map[string]string, error) {
	getSql := fmt.Sprintf("SELECT * FROM `%s` WHERE `%s` = ?", tableName, fieldName)
	return QueryOne(getSql, value)
}

func MustGetOneWhere(tableName, fieldName, value string) map[string]string {
	data, err := GetOneWhere(tableName, fieldName, value)
	if err != nil {
		panic(err)
	}
	return data
}

// GetAllInTable
// 获取表中全部数据
func GetAllInTable(tableName string) ([]map[string]string, error) {
	getSql := fmt.Sprintf("SELECT * FROM `%s`", tableName)
	return Query(getSql)
}

func MustGetAllInTable(tableName string) []map[string]string {
	data, err := GetAllInTable(tableName)
	if err != nil {
		panic(err)
	}
	return data
}

// RunSelectCommand
// 执行一条selectCommand
func RunSelectCommand(selectCommand *MysqlAst.SelectCommand) (mapValue []map[string]string, err error) {
	prepareSql, parameterList := selectCommand.GetPrepareParameter()
	mapValue, err = Query(prepareSql, argsStringToInterface(parameterList...)...)
	return
}

func MustRunSelectCommand(selectCommand *MysqlAst.SelectCommand) (mapValue []map[string]string) {
	mapValue, err := RunSelectCommand(selectCommand)
	if err != nil {
		panic(err)
	}
	return
}

// IsExist
// 根据传入的map判断表中是否存在
func IsExist(tableName string, row map[string]string) bool {
	where := MysqlAst.NewAndWhereCondition()
	for k, v := range row {
		where = where.AddPrepare(fmt.Sprintf("%s=?", k), v)
	}
	selectCommand := MysqlAst.NewSelectCommand().From(tableName).WhereObj(where)
	info, err := RunSelectCommand(selectCommand)
	if err != nil || len(info) == 0 {
		return false
	}
	return true
}
