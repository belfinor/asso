package words

// @author  Mikhail Kirillov <mikkirillov@yandx.ru>
// @version 1.000
// @date    2019-03-06

import (
	"net/http"

	"github.com/belfinor/Helium/net/http/params"
	"github.com/belfinor/Helium/net/http/router"
	"github.com/belfinor/Helium/tmpl"
	"github.com/belfinor/asso/assodb"
	"github.com/belfinor/asso/text"
)

func init() {
	router.Register("GET", "/asso/word", Handler)
}

var addTmpl = `<!DOCTYPE html>
<html>
<head>
  <title>Добавить слово</title>
</head>
<body>
<form method="GET">
<input type="text" name="name">
<input type="submit" value="Добавить">
</form>
</body>
</html>
`

func Handler(rw http.ResponseWriter, r *http.Request, prms router.Params) {

	args := params.New(r)

	wrd := args.GetString("name")

	wrd = text.Prepare(wrd)

	if wrd != "" {
		assodb.AddWord(wrd)
	}

	jet := tmpl.New(nil, false)

	jet.LoadTmplString("page.jet", addTmpl)

	vars := jet.Vars()
	text, _ := jet.Render("page.jet", vars)

	rw.Header().Set("Content-Type", "text/html; charset=utf-8")

	rw.WriteHeader(200)
	rw.Write([]byte(text))
}
