package cache

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-02-06

import (
	"github.com/belfinor/lcache"
)

var cache *lcache.Cache = lcache.New(&lcache.Config{Size: 102400, Nodes: 16, TTL: 86400 * 7})

func makeKey(game_id, word string) string {
	return game_id + "#" + word
}

func SaveWord(game_id, word string) {
	cache.Set(makeKey(game_id, word), true)
}

func HasWord(game_id, word string) bool {

	key := makeKey(game_id, word)

	res := cache.Get(key)

	return res != nil
}

func reset() {
	cache = lcache.New(&lcache.Config{Size: 100000, Nodes: 16, TTL: 86400})
}
