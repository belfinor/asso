package db

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-02-05

import (
	"database/sql"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/asso/config"

	_ "github.com/lib/pq"
)

func Open() (*sql.DB, error) {

	cfg := config.Get().Database

	db, err := sql.Open(cfg.Driver, cfg.Connect)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		db.Close()
		return nil, err
	}

	return db, err
}
