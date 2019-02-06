package avatars

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-02-06

import (
	"errors"
	"io/ioutil"
	"strings"
	"time"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/rand/mersenne"
	"github.com/belfinor/asso/config"
)

var data []string = []string{}

func init() {
	go reload()
}

func reload() {
	for {
		<-time.After(time.Minute * 10)
		Load()
	}
}

func Load() {

	cfg := config.Get()

	files, err := ioutil.ReadDir(cfg.Avatars.Path)
	if err != nil {
		log.Error(err)
		return
	}

	res := make([]string, 0, len(files))

	allow := map[string]bool{
		"gif":  true,
		"jpeg": true,
		"jpg":  true,
		"png":  true,
	}

	for _, f := range files {

		if f.IsDir() {
			continue
		}

		lst := strings.Split(f.Name(), ".")
		if len(lst) != 2 {
			continue
		}

		if _, has := allow[lst[1]]; !has {
			continue
		}

		res = append(res, f.Name())
	}

	data = res

	log.Info("avatars loaded")
}

func Size() int {

	return len(data)
}

func Get(index int) (string, error) {

	if index < 0 || index >= Size() {
		return "", errors.New("invalid index")
	}

	return config.Get().Avatars.Prefix + "/" + data[index], nil
}

func Random() (string, error) {

	if Size() < 1 {
		return "", errors.New("no data")
	}

	index := int(mersenne.Next() % int64(len(data)))

	return Get(index)
}
