package main

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2019-02-07

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/belfinor/Helium/daemon"
	"github.com/belfinor/Helium/log"
	"github.com/belfinor/asso/assodb"
	"github.com/belfinor/asso/config"
	"github.com/belfinor/asso/text"
)

func main() {
	conf := ""
	isDaemon := false
	isCon := false

	flag.StringVar(&conf, "c", "/etc/asso.json", "config file name")
	flag.BoolVar(&isDaemon, "d", false, "run as daemon")
	flag.BoolVar(&isCon, "console", false, "run in console mode")

	flag.Parse()

	cfg := config.Init(conf)

	if isDaemon {
		daemon.Run(&cfg.Daemon)
	}

	log.Init(&cfg.Log)

	if isDaemon {
		log.Info("start application as daemon")
	} else {
		log.Info("start application")
	}

	assodb.Init()

	if isCon {
		consoleHandler()
		log.Finish("finish")
		return
	}

	log.Finish("finish")
}

func consoleHandler() {

	br := bufio.NewReader(os.Stdin)

	start := assodb.Word()

	fmt.Println("start word: " + start)

	for {

		str, err := br.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		str = text.Prepare(str)
		if str == "" {
			continue
		}

		assodb.Add(start, str)

		start = str
	}
}
