package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"blogs-service/cmd/api/v1"
	"blogs-service/common"
	"blogs-service/config"
	"blogs-service/service/blogs"

	middleware "github.com/LYY/echo-middleware"
	"github.com/alecthomas/kingpin"
	"github.com/labstack/echo"
	emw "github.com/labstack/echo/middleware"
	"github.com/robfig/cron"
	validator "gopkg.in/go-playground/validator.v9"
)

var (
	// Version version
	Version = "0.0.1"
	app     = kingpin.New("app", "blogsapi applicaton server").DefaultEnvars()
	cmdRun  = app.Command("run", "Run application").Default()
)

func main() {
	kingpin.Version(Version)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	//init config
	config.Server = "blogs"
	conf := config.InitConfig()

	e := echo.New()
	//echo init ,include log init
	common.EchoInit(e, conf.Logpath, conf.Debug)
	e.Validator = &common.SimpleValidator{Validator: validator.New()}

	// middlewares
	e.Pre(emw.RemoveTrailingSlash())
	e.Pre(middleware.NoCache())
	e.Pre(middleware.Heartbeat("/ping"))
	e.Use(middleware.RequestID())
	e.Use(emw.Secure())
	e.Use(common.Recover())

	//InitDB
	common.InitDB(conf.Db.Mysqlurl, conf.Db.Mysqlidle, conf.Db.Mysqlmaxopen, conf.Debug)
	defer common.DBClose()
	blogs.DbMigrate()

	// register api
	v1.RegisterAPI(e)

	srvAddr := ":" + strconv.Itoa(conf.Port)

	e.Logger.Infof("Listening and serving HTTP on %s", srvAddr)
	e.Logger.Infof("debug: %s", strconv.FormatBool(conf.Debug))

	// Start server
	go func() {
		if err := e.Start(srvAddr); err != nil {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	//new and start cron
	c := cron.New()
	c.AddFunc("@every 1m", common.UpLogFile)
	c.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	//stop cron
	c.Stop()
	e.Logger.Info("Server exist")
}
