package assodb

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
// @date    2019-02-09

import (
	"context"
	"database/sql"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/asso/db"
	"github.com/belfinor/lcache"
)

var acache *lcache.Cache = lcache.New(&lcache.Config{TTL: 86400 * 7, Size: 10240, Nodes: 16})
var wstream chan string = make(chan string, 10)

func Word() string {
	return <-wstream
}

func Init() {

	go func() {
		for {
			for _, w := range getWords() {
				wstream <- w
			}
		}
	}()
}

func getWords() []string {
	dbh, err := db.Open()
	if err != nil {
		return nil
	}
	defer dbh.Close()

	rows, err := dbh.Query("SELECT name FROM words ORDER BY RANDOM() LIMIT 10")
	if err != nil {
		log.Error(err)
		return nil
	}

	res := make([]string, 0, 10)
	str := ""

	for rows.Next() {

		rows.Scan(&str)
		res = append(res, str)
	}

	return res
}

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
		log.Error(e)
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

		rows, e := dbh.Query("SELECT asso FROM asso WHERE name = $1 ORDER BY RANDOM() LIMIT 100", name)
		if e != nil {
			log.Error(e)
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
