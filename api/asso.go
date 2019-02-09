package api

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-02-09

import (
	"github.com/belfinor/Helium/net/jsonrpc2"
	"github.com/belfinor/asso/avatars"
	"github.com/belfinor/asso/errors"
	"github.com/belfinor/asso/game"
	"github.com/belfinor/asso/text"
)

func init() {
	jsonrpc2.RegisterMethod("asso.game", Game)
	jsonrpc2.RegisterMethod("asso.need_help", NeedHelp)
	jsonrpc2.RegisterMethod("asso.start_game", StartGame)
}

func StartGame(req map[string]string) (*game.Game, *jsonrpc2.Error) {

	wrd, has := req["start"]
	if has {
		wrd = text.Prepare(wrd)
	}

	return game.New(wrd), nil
}

func Game(req GameData) (*GameResponse, *jsonrpc2.Error) {
	word := text.Prepare(req.Word)
	word_before := text.Prepare(req.Before)

	if word == "" || word_before == "" {
		return nil, errors.Get("no_asso")
	}

	g := game.GetGame(req.GameId)

	if !g.UserStep(word_before, word) {
		return nil, errors.Get("word_already_used")
	}

	rec := g.ComputerStep(word)

	gr := &GameResponse{GameId: req.GameId}

	if rec == nil {
		gr.Victory = true
	} else {
		gr.Word = rec.Asso
		gr.Image = avatars.Get()
	}

	return gr, nil
}

func NeedHelp(req GameData) (*GameResponse, *jsonrpc2.Error) {
	word_before := text.Prepare(req.Before)

	if word_before == "" {
		return nil, errors.Get("no_asso")
	}

	g := game.GetGame(req.GameId)

	rec := g.ComputerStep(word_before)

	gr := &GameResponse{GameId: req.GameId}

	if rec == nil {
		gr.Victory = true
	} else {
		gr.Word = rec.Asso
		gr.Image = avatars.Get()
	}

	return gr, nil
}
