/**
 * Created by angelina on 2017/4/15.
 */

package yeeSql

import (
	"fmt"
)

// 注册的table表列表
var registerTableList = []Table{}

// MustRegisterTable
// 向系统中注册表
// 不能重复注册
// 不能并发调用
func MustRegisterTable(table Table) {
	for i := range registerTableList {
		if registerTableList[i].Name == table.Name {
			panic(fmt.Errorf("[MustRegisterTable] table name %s repeat", table.Name))
		}
	}
	registerTableList = append(registerTableList, table)
}

// RegisterTables
// 获取已注册的表
func RegisterTables() []Table {
	return registerTableList
}

// ClearRegisterTable
// 清除系统中注册的全部数据库表
// 不能并发调用
func ClearRegisterTable() {
	registerTableList = []Table{}
}

// MustSyncRegisterTable
// 同步注册进去的表(只会增加字段,保证不掉数据,会使用fmt显示有哪些字段存在问题.)
// 不能并发调用
func MustSyncRegisterTable() {
	for i := range registerTableList {
		MustSyncTable(registerTableList[i])
	}
}

// MustForceSyncRegisterTable
// 强制同步注册进去的表(可能会缺失字段，保证字段达到配置的样子)
// 不能并发调用
func MustForceSyncRegisterTable() {
	for i := range registerTableList {
		MustForceSyncTable(registerTableList[i])
	}
}

// MustCreateDb
// 根据传入的配置创建数据库DB
func MustCreateDb() {
	GetDbWithoutDbName().MustExec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s "+
		"DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_bin;", GetDbConfig().DbName))
}

func MustDropDb() {
	GetDbWithoutDbName().MustExec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", GetDbConfig().DbName))
}
