/**
 * Created by angelina on 2017/4/15.
 */

package yeeSql

import (
	"fmt"
	"github.com/vannnnish/yeego/yeeStrings"
	"sort"
	"strings"
)

//读取数据库的表字段,不区分大小写(某些系统的mysql不区分大小写)
//写入数据库的表字段,区分大小写
type Table struct {
	Name       string
	FieldList  map[string]DbType
	PrimaryKey string
	UniqueKey  [][]string
	NotNull    []string
}

type DbType string

const (
	DbTypeInt              DbType = `int(11) DEFAULT 0`
	DbTypeIntAutoIncrement DbType = `int(11) unsigned AUTO_INCREMENT`
	DbTypeString           DbType = `varchar(255) COLLATE utf8mb4_bin DEFAULT ""`
	DbTypeLongString       DbType = `longtext COLLATE utf8mb4_bin DEFAULT ""`
	DbTypeFloat            DbType = `float default 0`
	DbTypeDatetime         DbType = `datetime DEFAULT "1970-01-01 00:08:00"`
	DbTypeBool             DbType = `tinyint(4) DEFAULT 0`
	DbTypeLongBlob         DbType = `LONGBLOB`
)

func (t DbType) GetMysqlFieldType() MysqlFieldType {
	switch t {
	case DbTypeInt:
		return MysqlFieldType{
			DataType: MysqlDataTypeInt32,
			Default:  "0",
		}
	case DbTypeIntAutoIncrement:
		return MysqlFieldType{
			DataType:        MysqlDataTypeInt32,
			IsUnsigned:      true,
			IsAutoIncrement: true,
		}
	case DbTypeString:
		return MysqlFieldType{
			DataType:         MysqlDataTypeVarchar,
			Default:          "",
			CharacterSetName: "utf8mb4",
			CollationName:    "utf8mb4_bin",
			StringLength:     255,
		}
	case DbTypeLongString:
		return MysqlFieldType{
			DataType:         MysqlDataTypeLongText,
			Default:          "",
			CharacterSetName: "utf8mb4",
			CollationName:    "utf8mb4_bin",
		}
	case DbTypeFloat:
		return MysqlFieldType{
			DataType: MysqlDataTypeFloat,
			Default:  "0",
		}
	case DbTypeDatetime:
		return MysqlFieldType{
			DataType: MysqlDataTypeDateTime,
			Default:  "1970-01-01 00:08:00",
		}
	case DbTypeBool:
		return MysqlFieldType{
			DataType: MysqlDataTypeInt8,
			Default:  "0",
		}
	case DbTypeLongBlob:
		return MysqlFieldType{
			DataType: MysqlDataTypeLongBlob,
		}
	default:
		panic(fmt.Errorf("Unsupport DbType %s", t))
	}
}

// MustVerifyTableConfig
// 检测table的字段正确性，不能有重复字段(忽略大小写)
func MustVerifyTableConfig(table Table) {
	fieldNameMap := map[string]bool{}
	for name := range table.FieldList {
		name = strings.ToLower(name)
		if fieldNameMap[name] {
			panic(fmt.Errorf("[MustVerifyTableConfig] Table[%s] Field[%s] 两个字段名只有大小写不一致",
				table.Name, name))
		}
		fieldNameMap[name] = true
	}
}

// MustIsTableExist
// 检测这个table是否存在
func MustIsTableExist(tableName string) bool {
	ret, err := QueryOne("SHOW TABLE STATUS WHERE Name=?", tableName)
	if err != nil {
		return false
	}
	if len(ret) <= 0 {
		return false
	} else {
		return true
	}
}

// MustCreateTable
// 根据table建表
/*
	CREATE TABLE IF NOT EXISTS `test`
	(
		ID int(11) unsigned AUTO_INCREMENT,
		Name varchar(255) COLLATE utf8_bin DEFAULT "" NOT NULL,
		Account varchar(255) COLLATE utf8_bin DEFAULT "" NOT NULL,
		PRIMARY KEY (`ID`),
		UNIQUE INDEX (`Account`,`Name`)
	) engine=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin
*/
func MustCreateTable(table Table) {
	sql := "CREATE TABLE IF NOT EXISTS `" + table.Name + "` \n("
	sqlItemList := []string{}
	hasPrimaryKey := false
	if v, ok := table.FieldList[table.PrimaryKey]; ok {
		sqlField := "`" + table.PrimaryKey + "` " + string(v)
		sqlItemList = append(sqlItemList, sqlField)
	}
	keys := make([]string, 0)
	for k := range table.FieldList {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fieldName := k
		fieldType := table.FieldList[k]
		if table.PrimaryKey == fieldName {
			hasPrimaryKey = true
			continue
		}
		sqlField := "`" + fieldName + "` " + string(fieldType)
		if yeeStrings.IsInSlice(table.NotNull, fieldName) {
			sqlField += " NOT NULL"
		}
		sqlItemList = append(sqlItemList, sqlField)
	}
	//for fieldName, fieldType := range table.FieldList {
	//	if table.PrimaryKey == fieldName {
	//		hasPrimaryKey = true
	//		continue
	//	}
	//	sqlField := "`" + fieldName + "` " + string(fieldType)
	//	if yeeStrings.IsInSlice(table.NotNull, fieldName) {
	//		sqlField += " NOT NULL"
	//	}
	//	sqlItemList = append(sqlItemList, sqlField)
	//}
	if table.PrimaryKey != "" {
		if !hasPrimaryKey {
			panic(fmt.Sprintf(`table.PrimaryKey[%s], 但是这个主键不在字段列表里面`, table.PrimaryKey))
		}
		sqlItemList = append(sqlItemList, "PRIMARY KEY (`"+table.PrimaryKey+"`)")
	}
	for _, group := range table.UniqueKey {
		uniqueSql := "UNIQUE INDEX ("
		uniqueKeyList := []string{}
		for _, key := range group {
			uniqueKeyList = append(uniqueKeyList, "`"+key+"`")
		}
		uniqueSql += strings.Join(uniqueKeyList, ",") + ")"
		sqlItemList = append(sqlItemList, uniqueSql)
	}
	sql += strings.Join(sqlItemList, ",\n")
	sql += "\n) engine=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin"
	MustExec(sql)
}

// MustAddField
// 添加字段
func MustAddField(table Table, filedName string) {
	newFieldType := table.FieldList[filedName]
	sql := "ALTER TABLE `" + table.Name + "` ADD `" + filedName + "` " + string(newFieldType)
	if yeeStrings.IsInSlice(table.NotNull, filedName) {
		sql += " NOT NULL"
	}
	MustExec(sql)
}

// MustModifyTable
// 更新表数据
func MustModifyTable(table Table) {
	// 获取真实数据库中的全部字段
	MysqlFieldTypeList := mustMysqlGetTableFieldTypeList(table.Name)
	// 数据库表字段名称
	dbFieldNameList := []string{}
	for _, row := range MysqlFieldTypeList {
		dbFieldNameList = append(dbFieldNameList, strings.ToLower(row.Name))
	}
	for _, f1 := range dbFieldNameList {
		found := false
		for f2 := range table.FieldList {
			if strings.EqualFold(f1, f2) {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("[yeeSql.SyncTable] 数据库中包含多余字段 Table[%s] Field[%s]\n", table.Name, f1)
		}
	}
	for fieldName, fieldType := range table.FieldList {
		if !yeeStrings.IsInSlice(dbFieldNameList, strings.ToLower(fieldName)) {
			MustAddField(table, fieldName)
			continue
		}
		for _, row := range MysqlFieldTypeList {
			if fieldName == row.Name {
				// 相同字段名称的类型不相同
				if !fieldType.GetMysqlFieldType().Equal(row.Type) {
					fmt.Printf("[yeeSql.SyncTable] Table[%s] Field[%s] OldType[%s] NewType[%s] 数据库字段类型不一致\n",
						table.Name, fieldName, row.Type.String(), fieldType.GetMysqlFieldType().String())
				}
				break
			}
			if strings.EqualFold(row.Name, fieldName) {
				fmt.Printf("[yeeSql.SyncTable] Table[%s] OldField[%s] NewField[%s] 数据库字段大小写不一致\n",
					table.Name, fieldName, row.Name)
				break
			}
		}
	}
}

// MustForceModifyTable
// 强制更新
func MustForceModifyTable(table Table) {
	MysqlFieldTypeList := mustMysqlGetTableFieldTypeList(table.Name)
	dbFieldNameList := []string{}
	for _, row := range MysqlFieldTypeList {
		dbFieldNameList = append(dbFieldNameList, row.Name)
	}
	for _, f1 := range dbFieldNameList {
		found := false
		for f2 := range table.FieldList {
			if f2 == f1 {
				found = true
				break
			}
		}
		if !found {
			MustExec(fmt.Sprintf("ALTER TABLE `%s` DROP COLUMN `%s`", table.Name, f1))
		}
	}
	for fieldName, fieldType := range table.FieldList {
		if yeeStrings.IsInSlice(dbFieldNameList, fieldName) {
			for _, row := range MysqlFieldTypeList {
				if row.Name == fieldName {
					if !fieldType.GetMysqlFieldType().Equal(row.Type) {
						MustExec(fmt.Sprintf("ALTER TABLE `%s` CHANGE COLUMN `%s` `%s` %s NOT NULL",
							table.Name, fieldName, fieldName, fieldType))
					}
					break
				}
			}
			continue
		}
		MustAddField(table, fieldName)
	}
}

// MustDropTable
// 删除表
func MustDropTable(tableName string) {
	MustExec("DROP TABLE IF EXIST `" + tableName + "`")
}

// MustSyncTable
// 同步更新表数据
// 只会增加字段,保证不掉数据,会使用fmt显示有哪些字段存在问题
func MustSyncTable(table Table) {
	MustVerifyTableConfig(table)
	if MustIsTableExist(table.Name) {
		MustModifyTable(table)
	} else {
		MustCreateTable(table)
	}
}

// MustForceSyncTable
// 强制更新表
// 可能会缺失字段，保证字段达到配置的样子
func MustForceSyncTable(table Table) {
	MustVerifyTableConfig(table)
	if MustIsTableExist(table.Name) {
		MustForceModifyTable(table)
	} else {
		MustCreateTable(table)
	}
}
