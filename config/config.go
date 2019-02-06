package config

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-02-06

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/belfinor/Helium/daemon"
	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/net/jsonrpc2"
)

type Config struct {
	Avatars struct {
		Path   string `json:"path"`
		Prefix string `json:"prefix"`
	} `json:"avatars"`
	Daemon   daemon.Config `json:"daemon"`
	Database struct {
		Driver  string `json:"driver"`
		Connect string `json:"connect"`
	} `json:"database"`
	Server jsonrpc2.HttpConfig `json:"server"`
	Log    log.Config          `json:"log"`
	GC     int                 `json:"gc"`
}

var _config *Config

func Load(filename string) *Config {

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(fmt.Sprintf("Open file %s error", filename))
	}

	var _con Config

	err = json.Unmarshal(data, &_con)

	if err != nil {
		panic(err)
	}

	return &_con
}

func Set(conf *Config) {
	_config = conf
}

func Init(conf string) *Config {
	if _config == nil {
		_config = Load(conf)
	}

	return _config
}

func Get() *Config {
	return _config
}
