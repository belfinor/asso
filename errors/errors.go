package errors

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-02-08

import (
	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/net/jsonrpc2"
)

// next -33006

var data map[string]*jsonrpc2.Error = map[string]*jsonrpc2.Error{
	"internal_error":    &jsonrpc2.Error{Code: -32603, Message: "Internal error"},
	"no_source":         &jsonrpc2.Error{Code: -33000, Message: "Текущая фраза не задана"},
	"no_cat":            &jsonrpc2.Error{Code: -33001, Message: "Категория не задана"},
	"no_word":           &jsonrpc2.Error{Code: -33002, Message: "Слово не найдено"},
	"no_asso":           &jsonrpc2.Error{Code: -33003, Message: "Ассоциация не задана"},
	"asso_already_used": &jsonrpc2.Error{Code: -33004, Message: "Ассоциация уже была использована"},
	"word_already_used": &jsonrpc2.Error{Code: -33005, Message: "Слово уже было использовано"},
}

func Get(code string) *jsonrpc2.Error {
	e, has := data[code]
	if has {
		return e
	}

	log.Error("unknown error code \"" + code + "\"")

	return &jsonrpc2.Error{Code: -32603, Message: "Internal error"}
}
