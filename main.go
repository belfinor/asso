package main

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.005
// @date    2019-03-06

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/belfinor/Helium/daemon"
	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/net/http/router"
	"github.com/belfinor/Helium/net/jsonrpc2"
	"github.com/belfinor/asso/assodb"
	"github.com/belfinor/asso/avatars"
	"github.com/belfinor/asso/config"
	"github.com/belfinor/asso/text"

	_ "github.com/belfinor/asso/api"

	_ "github.com/belfinor/asso/controller"
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

	avatars.Init()

	// start server
	router.Register("POST", cfg.Server.Url, func(rw http.ResponseWriter, req *http.Request, p router.Params) {
		jsonrpc2.HttpHandler(rw, req)
	})

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	router.GetDefault().LoggerFunc = log.Info

	log.Info("start http server " + addr)

	http.ListenAndServe(addr, router.GetDefault())

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
