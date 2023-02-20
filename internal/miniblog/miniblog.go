package miniblog

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"github.com/sjxiang/miniblog/internal/pkg/log"
	mw "github.com/sjxiang/miniblog/internal/pkg/middleware"
)

var (
	cfgFile string
	env     *Env
)

// NewMiniBlogCommand 创建一个 *cobra.Command 对象，之后，可以使用 Command 对象的 Execute 方法来启动应用程序
func NewMiniBlogCommand() *cobra.Command {

	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "miniblog",
		// 命令的简短描述
		Short: "迷你博客",
		// 命令的详细描述
		Long: `A good Go practical project，迷你博客。

Find more miniblog infomation at:
		https://github.com/sjxiang/miniblog#readme`,
		// 命令出错时，不打印帮助信息。（不需要，设置为 true，可以保持命令出错时一眼就能看到错误信息）
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {

			// 加载配置
			env = NewEnv(cfgFile)

			// 初始化日志
			log.Init(logOptions())
			defer log.Sync() // Sync 将缓存中的日志刷新到磁盘文件中

			return run()
		},
		// 内置`验证`函数
		// Args: cobra.MinimumNArgs(1),

		// 自定义`验证`函数（不需要指定命令行参数）
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	// 正常：Execute() -> Run()
	// 加塞：Execute() -> init() -> Run()
	// cobra.OnInitialize()

	// 在这里您将定义标志和配置设置

	// Cobra 支持持久性标志（PersistentFlags），该标志可用于它所分配的命令以及该命令下的每个子命令

	// 定义标准 Flag
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "miniblog 的配置文件路径，若字符串为空，则为无配置文件。")

	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	// cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}

// run 实际的业务代码入口
func run() error {

	log.Infow("Using config file", "file", env)

	// 设置 Gin 模式
	gin.SetMode(env.RunMode)

	g := gin.New()
	mws := []gin.HandlerFunc{
		gin.Recovery(), // 从任何一个 panic 中恢复，并返回 500
		mw.RequestID(),
		mw.Cors(),
	}
	g.Use(mws...)

	// 注册 404
	g.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Code":    10003,
			"message": "Page not found",
		})
	})

	// 健康检查
	g.GET("/healthz", func(ctx *gin.Context) {
		log.C(ctx).Infow("健康检查调用")

		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	// 创建 HTTP Server 实例
	httpsrv := &http.Server{
		Addr:    env.Addr,
		Handler: g,
	}

	// 运行 HTTP 服务器
	log.Infow("开始监听 HTTP 地址上的请求", "addr", env.Addr)
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()

	// 开始监听系统信号（优雅关停）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, signals...)
	<-quit
	log.Infow("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}

	log.Infow("Server exiting")
	return nil
}
