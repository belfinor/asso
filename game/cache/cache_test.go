package cache

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-02-06

import (
	"testing"
)

func TestCache(t *testing.T) {
	reset()

	if HasWord("1", "2") {
		t.Fatal("cache not empty")
	}

	SaveWord("1", "2")

	if !HasWord("1", "2") {
		t.Fatal("SetWord not work")
	}

	if HasWord("2", "3") {
		t.Fatal("Has word fount not existing key")
	}
}
