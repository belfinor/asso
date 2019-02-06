package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-02-06

import (
	"strings"

	"github.com/belfinor/Helium/text"
)

func Prepare(str string) string {

	list := text.GetWords(str)
	res := strings.Join(list, " ")

	return res

}
