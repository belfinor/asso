package avatars

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2019-04-17

import (
	"io/ioutil"
	"strings"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/asso/config"
	"github.com/belfinor/lcache"
	"github.com/belfinor/lrand"
)

var data []string = []string{}
var stream chan string = make(chan string, 20)
var atom *lcache.Atom = lcache.NewAtom(600)

func Init() {

	go func() {

		for {

			atom.Fetch(func() interface{} {
				load()
				return true
			})

			index := int(lrand.Next() % int64(Size()))
			stream <- config.Get().Avatars.Prefix + "/" + data[index]

		}

	}()
}

func load() {

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

	log.Info("avatars reloaded")
}

func Size() int {
	return len(data)
}

func Get() string {
	return <-stream
}
