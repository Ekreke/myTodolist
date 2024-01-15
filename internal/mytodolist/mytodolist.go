package mytodolist

import (
	"fmt"

	"github.com/ekreke/myTodolist/internal/pkg/log"
	mw "github.com/ekreke/myTodolist/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// newmytodolist 创建一个 *cobra.Command 对象之后 ， 可以使用Command对象的execute方法来启动一个应用程序
func NewMyTodolistCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "mytodolist",
		// 指定命令的简短描述信息，该描述信息会出现在帮助信息中
		Short: "mytodolist is a todo list application",
		// 指令的详细描述
		Long: "mytodolist is a todo list application , used to create a todolist server",

		SilenceUsage: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			log.Init(logOptions())
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}
	// 以下设置，使得 initConfig 函数在每个命令运行时都会被调用以读取配置
	cobra.OnInitialize(initConfig)

	// 在这里您将定义标志和配置设置。

	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the mytodolist configuration file. Empty string for no configuration file.")

	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// 添加 --version 标志
	// verflag.AddFlags(cmd.PersistentFlags())

	return cmd
}

func run() error {
	g := gin.New()
	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.Secure, mw.RequestID()}
	g.Use(mws...)
	// settings, _ := json.Marshal(viper.AllSettings())
	// log.Infow("settings", "settings", string(settings))
	log.Infow(viper.GetString("db.username"))
	// 打印 db -> username 配置项的值
	return nil
}
