package main

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-02-06

import (
	"flag"

	"github.com/belfinor/Helium/daemon"
	"github.com/belfinor/Helium/log"
	"github.com/belfinor/asso/config"
)

func main() {
	conf := ""
	is_daemon := false

	flag.StringVar(&conf, "c", "/etc/asso.json", "config file name")
	flag.BoolVar(&is_daemon, "d", false, "run as daemon")

	flag.Parse()

	cfg := config.Init(conf)

	if is_daemon {
		daemon.Run(&cfg.Daemon)
	}

	log.Init(&cfg.Log)

	if is_daemon {
		log.Info("start application as daemon")
	} else {
		log.Info("start application")
	}

}
