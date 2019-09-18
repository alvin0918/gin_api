package run

import (
	"fmt"
	"github.com/alvin0918/gin_api/core/config"
	"github.com/alvin0918/gin_api/core/Initialization"
	"github.com/alvin0918/gin_api/core/middleware"
	"github.com/alvin0918/gin_api/routers"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func Run() {

	var (
		err error
	)

	// 解析命令行参数
	Initialization.InitArgs()

	// 初始化线程数量
	Initialization.InitEnv()

	// 加载配置
	if err = config.InitConfig(Initialization.ConfFile); err != nil {
		log.Fatal(err)
	}

	// 启动成功记录PID
	Initialization.InitPid()

	// 启动GIN服务
	runGin()
}

func runGin() {
	var (
		r      *gin.Engine
		srv    *http.Server
		err    error
		str    os.Signal
		ctx    context.Context
		cancel context.CancelFunc
		quit   chan os.Signal
	)

	// 注册中间件
	r = gin.New()

	// 加载必要middleware
	r.Use(middleware.AccessJsMiddleware())
	r.Use(middleware.ErrorMiddleware())

	// 加载路由
	routers.Init(r)

	// 启动服务
	printColor("开始运行：http://" + config.G_config.ApiHostAndPort)
	srv = &http.Server{
		Addr:    config.G_config.ApiHostAndPort,
		Handler: r,
	}

	// 启用协程链接服务
	go func() {
		// 服务链接
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// 出现错误，则关闭服务
			exec.Command("kill", "-s SIGINT ", string(os.Getppid()))
		}

	}()

	// 等待中断信号以优雅地关闭服务器（设置 3 秒的超时时间）
	quit = make(chan os.Signal)

	// 监听信号
	signal.Notify(quit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1)

	// 监听信道信号
	str = <-quit

	switch str {

	case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt:
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		// 关闭通到
		defer cancel()

		if err = srv.Shutdown(ctx); err != nil {
			printColor(fmt.Sprintf("Server Shutdown %s", err.Error()))
		}

		// 删除PID文件
		printColor("删除PID文件")
		os.Remove(".pid")

		printColor("Server exiting. Bey!")

	case syscall.SIGUSR1:

		printColor(fmt.Sprintf("[%s]Closing Server...", str))

		// 保存进程上下文， 在运行3秒
		printColor("剩下任务处理")
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		// 关闭通到
		defer cancel()

		printColor("开始关闭Server")
		if err = srv.Shutdown(ctx); err != nil {
			printColor(fmt.Sprintf("Server Shutdown %s", err.Error()))
		}

		// 删除PID文件
		printColor("删除PID文件")
		os.Remove(".pid")

		printColor("开始重启")

		Run()

	default:
		printColor("参数错误,可使用-h查看帮助")
	}

}

func printColor(str interface{}) {
	var (
		cyan  = string([]byte{27, 91, 51, 54, 109})
		reset = string([]byte{27, 91, 48, 109})
	)

	fmt.Println(cyan, str, reset)
}
