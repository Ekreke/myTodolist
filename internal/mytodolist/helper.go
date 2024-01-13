package mytodolist

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// recommendedHomeDir 定义放置服务配置的默认目录 maybe linux？
	recommendedHomeDir = ".myTodolist"
	// defaultConfigName 指定了默认配置文件名称
	defaultConfigName = "mytodolist"
)

// initconfig 设置需要读取的配置文件名、环境变量、并读取配置文件到viper中
// 支持 环境变量/ 命令行选定路径 /开发版本默认路径
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(defaultConfigName)
	}
	viper.AutomaticEnv()
	// 读取环境变量的前缀为 mytodolist，如果是 mytodolist，将自动转变为大写。
	viper.SetEnvPrefix("MYTODOLIST")
	// 以下 2 行，将 viper.Get(key) key 字符串中 '.' 和 '-' 替换为 '_'
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	// 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		if err != nil {
			// 开发版 配置文件保存在configs目录下
			viper.SetConfigFile("../../configs/mytodolist.yaml")
			fmt.Println("111")
		}
		// log.Errorw("Failed to read viper configuration file", "err", err)
	}

	if err := viper.ReadInConfig(); err != nil {
		if err != nil {
			fmt.Println("111")

		}
	}
	// 打印 viper 当前使用的配置文件，方便 Debug.
	// log.Debugw("Using config file", "file", viper.ConfigFileUsed())
}

// logOptions 从 viper 中读取日志配置，构建 `*log.Options` 并返回.
// 注意：`viper.Get<Type>()` 中 key 的名字需要使用 `.` 分割，以跟 YAML 中保持相同的缩进.
func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}
