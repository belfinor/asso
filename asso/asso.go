package asso

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-02-06

import (
	"context"
	"database/sql"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/asso/db"
	"github.com/belfinor/lcache"
)

var acache *lcache.Cache = lcache.New(&lcache.Config{TTL: 86400 * 7, Size: 10240, Nodes: 16})

func Add(from, to string) {

	dbh, err := db.Open()
	if err != nil {
		return
	}
	defer dbh.Close()

	tx, e := dbh.BeginTx(context.Background(), nil)
	if e != nil {
		log.Error(e)
		return
	}
	defer tx.Rollback()

	row := tx.QueryRow("SELECT id FROM asso WHERE name = $1 AND asso = $2", from, to)

	var id int

	err = row.Scan(&id)

	if err == sql.ErrNoRows {

	} else if err != nil {
		return
	}

	if id == 0 {
		tx.Exec("INSERT INTO asso(name,asso) VALUES ($1,$2)", from, to)
		acache.Delete(from)
	} else {
		tx.Exec("UPDATE asso SET counter = counter + 1 WHERE id = $1", id)
	}

	tx.Commit()
}

func Asso(name string) []string {

	result := acache.Fetch(name, func(name string) interface{} {

		dbh, err := db.Open()
		if err != nil {
			return []string{}
		}
		defer dbh.Close()

		rows, e := dbh.Query("SELECT asso.asso FROM asso WHERE name = $1 ORDER BY RANDOM(1000) LIMIT 100", name)
		if e != nil {
			return []string{}
		}

		res := []string{}
		var str string

		for rows.Next() {
			rows.Scan(&str)
			res = append(res, str)
		}

		return res
	})

	return result.([]string)
}

func Delete(from, to string) {

	dbh, err := db.Open()
	if err != nil {
		return
	}
	defer dbh.Close()

	dbh.Exec("DELETE FROM asso WHERE name = $1 AND asso = $2", from, to)

	acache.Delete(from)
}
