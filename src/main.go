package main

import (
	"logger"
	"github.com/op/go-logging"
	"config"
	"github.com/go-martini/martini"
	"routing"
	"db"
	"github.com/martini-contrib/render"
)

func main() {
	var (
		err error
		cmdConfig *config.CommandLineConfiguration
		cfg *config.Config
		m *martini.ClassicMartini
		log *logging.Logger
	)
	cmdConfig, err = config.GetCommandLineConfiguration()
	if err != nil {
		panic(err)
	}
	cfg, err = config.GetConfig(cmdConfig.ConfigFilePath)
	if err != nil {
		panic(err)
	}
	defer cfg.Logger.LogFile.Close()
	log, err = logger.GetLogger(&cfg.Logger)
	if err != nil {
		panic(err)
	}
	session, err := db.GetSession(&cfg.DatabaseConnectOpts)
	if err != nil {
		panic(err)
	}
	m = martini.Classic()
	m.Map(log)
	m.Map(&cfg.WebSocketConfig)
	m.Map(&cfg.HttpServer)
	m.Map(&cfg)
	m.Map(session)
	m.Use(render.Renderer(render.Options{
		Layout: "base",
	}))
	routing.Configure(m)
	log.Info("Listening")
	m.RunOnAddr(cfg.ListenHost + ":" + cfg.ListenPort)
	log.Info("Listening")
}
