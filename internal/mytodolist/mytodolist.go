package mytodolist

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ekreke/myTodolist/internal/pkg/known"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	mw "github.com/ekreke/myTodolist/internal/pkg/middleware"
	"github.com/ekreke/myTodolist/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// newmytodolist 创建一个 *cobra.Command 对象之后 ， 可以使用Command对象的execute方法来启动一个应用程序
func NewMyTodolistCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "mytodolist",
		Short:        "mytodolist is a todo list application",
		Long:         "mytodolist is a todo list application , used to create a todolist server",
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
	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the mytodolist configuration file. Empty string for no configuration file.")
	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// 添加 --version 标志
	// verflag.AddFlags(cmd.PersistentFlags())
	return cmd
}

// run函数是实际的业务代码的入口函数
func run() error {

	if err := initStore(); err != nil {
		return err
	}
	// read jwt secret from config to signed token && parase
	token.Init(viper.GetString("jwt-sercret"), known.XUsernameKey)
	// set gin mode
	gin.SetMode(viper.GetString("runmode"))
	// create gin engine
	g := gin.New()
	// use gin.recovery / no cache / cors / secure / requestID middleware
	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.Secure, mw.RequestID()}
	g.Use(mws...)
	// install router
	if err := installRouters(g); err != nil {
		return err
	}
	g.GET("/test2", mw.Authn())
	// start a insecure server
	httpsrv := startInsecureServer(g)
	//  graceful quit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infow("Shutting down server...")
	// wait for the server handle remaining svc
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}
	log.Infow("Server exiting")
	return nil
}

// 创建并运行http服务器
func startInsecureServer(g *gin.Engine) *http.Server {
	// 创建 HTTP Server 实例
	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}
	// 运行HTTP服务器。 在goroutine中启动服务器，它不会阻止下面的正常关闭处理流程
	log.Infow("start to listening the incoming requests on http address", "addr", viper.GetString("addr"))
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()
	return httpsrv
}

// TODO:startSecureServer
// test graceful quit
func lazy() string {
	time.Sleep(10 * time.Second)
	return "lazy"
}
