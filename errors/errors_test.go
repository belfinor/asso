package errors

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-02-08

import (
	"testing"
)

func TestGet(t *testing.T) {
	if Get("no_asso").Code != -33003 {
		t.Fatal("no_asso wrong code")
	}

	if Get("asfasdfsdfasdf").Code != -32603 {
		t.Fatal("unknown error value wrong code")
	}
}
