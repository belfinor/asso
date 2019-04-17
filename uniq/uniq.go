package uniq

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-04-17

import (
	u "github.com/belfinor/luniq"
)

var unique *u.Uniq = u.New()

func Get() string {
	return unique.Next()
}
