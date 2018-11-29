/**
 * Created by angelina on 2017/4/15.
 */

package yeeSql

import "fmt"

// DbConfig
// 数据库配置
type DbConfig struct {
	UserName string // 用户名 example: root
	Password string // 密码 example: password
	Host     string // 数据库主机地址 example: 127.0.0.1
	Port     string // 数据库端口 example: 3306
	DbName   string // 数据库名称 example: yee_test
}

// MustVerifyDbConfig
// 验证参数是否正确
func MustVerifyDbConfig() {
	if dbConfig.UserName == "" {
		panic("UserName 不能为空")
	}
	if dbConfig.Password == "" {
		panic("Password 不能为空")
	}
	if dbConfig.Host == "" {
		panic("Host 不能为空")
	}
	if dbConfig.Port == "" {
		panic("Port 不能为空")
	}
	if dbConfig.DbName == "" {
		panic("DbName 不能为空")
	}
}

func MustSetDbConfig(conf *DbConfig) {
	dbConfig = *conf
	MustVerifyDbConfig()
}

// GetDsn
func (config *DbConfig) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&timeout=10s",
		config.UserName,
		config.Password,
		config.Host,
		config.Port,
		config.DbName)
}

// GetDsnWithoutDbName
func (config *DbConfig) GetDsnWithoutDbName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&timeout=10s",
		config.UserName,
		config.Password,
		config.Host,
		config.Port)
}

func GetDbConfig() DbConfig {
	MustVerifyDbConfig()
	return dbConfig
}
