package uniq

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-02-06

import (
	"testing"
)

func TestGet(t *testing.T) {
	hash := make(map[string]bool)

	for i := 0; i < 1000; i++ {
		hash[unique.Next()] = true
	}

	if len(hash) != 1000 {
		t.Fatal("uniq not work")
	}
}
