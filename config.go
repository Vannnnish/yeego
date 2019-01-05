package yeego

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"yeego/yeeFile"
	"os"
	"path/filepath"
	"strings"
)

type configType struct {
	*viper.Viper
}

var Config configType

func MustInitConfig(filePath string, fileName string) {
	Config = configType{viper.New()}
	Config.SetConfigName(fileName)
	//filePath支持相对路径和绝对路径 etc:"/a/b" "b" "./b"
	if filePath[:1] != "/" {
		// 相对路径
		Config.AddConfigPath(GetCurrentPath(filePath) + "/")
	} else {
		// 绝对路径
		Config.AddConfigPath(filePath + "/")
	}
	if err := Config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err).Error())
	}
}

func WatchConfig() {
	Config.WatchConfig()
	Config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}

var WORK_PATH string

// GetCurrentPath
// 获取项目路径下面的一些目录，不存在直接panic
func GetCurrentPath(dirPath string) string {
	if WORK_PATH != "" {
		return WORK_PATH
	}
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appDirPath := filepath.Join(appPath, dirPath)
	if !yeeFile.FileExists(appDirPath) {
		appDirPath = filepath.Join(workPath, dirPath)
		if !yeeFile.FileExists(appDirPath) {
			panic(fmt.Sprintf("dirPath:[%s] can not find in %s and %s", dirPath, appPath, workPath))
		}
	}
	WORK_PATH = strings.Replace(appDirPath, "\\", "/", -1)
	return WORK_PATH
}
