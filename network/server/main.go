package main

import (
	"BurninProject/aop/logger"
	"BurninProject/engine/world"
	"flag"
	"github.com/phuhao00/sugar"
	"github.com/xiaonanln/goworld/engine/config"
	"math/rand"
	"time"
)

var (
	gameid          uint16
	configFile      string
	logLevel        string
	restore         bool
	runInDaemonMode bool
	//signalChan      = make(chan os.Signal, 1)
)

func parseArgs() {
	var gameidArg int
	flag.IntVar(&gameidArg, "gid", 0, "set gameid")
	flag.StringVar(&configFile, "configfile", "", "set config file path")
	flag.StringVar(&logLevel, "log", "", "set log level, will override log level in config")
	flag.BoolVar(&restore, "restore", false, "restore from freezed state")
	flag.BoolVar(&runInDaemonMode, "d", false, "run in daemon mode")
	flag.Parse()
	gameid = uint16(gameidArg)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	parseArgs()

	if configFile != "" {
		config.SetConfigFile(configFile)
	}

	world.MM = world.NewMgrMgr()
	go world.MM.Run()
	logger.Logger.InfoF("server start !!")
	sugar.WaitSignal(world.MM.OnSystemSignal)
}
