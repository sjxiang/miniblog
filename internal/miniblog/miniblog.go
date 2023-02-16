package miniblog

import (
	"fmt"

	"github.com/spf13/cobra"
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
	// Flag
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "miniblog 的配置文件路径，若字符串为空，则为无配置文件。")

	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	// cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}

// run 实际的业务代码入口
func run() error {

	// env := NewEnv(cfgFile)
	fmt.Println(env.AppEnv)

	return nil
}
