package uniq

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-02-06

import (
	u "github.com/belfinor/Helium/uniq"
)

var unique *u.Uniq = u.New()

func Get() string {
	return unique.Next()
}
