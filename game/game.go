package game

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-02-06

import (
	"context"
	"time"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/slice"
	"github.com/belfinor/asso/asso"
	ass "github.com/belfinor/asso/asso"
	"github.com/belfinor/asso/avatars"
	"github.com/belfinor/asso/game/cache"
	"github.com/belfinor/asso/text"
	"github.com/belfinor/asso/uniq"
	"github.com/belfinor/sociation"
)

type Game struct {
	Id   string `json:"game_id"`
	Word string `json:"word"`
}

type Record struct {
	Word  string `json:"word"`
	Asso  string `json:"asso"`
	Image string `json:"image"`
}

func NextId() string {
	return uniq.Get()
}

func New(wrd string) *Game {
	g := &Game{
		Id: NextId(),
	}

	if wrd == "" {
		g.Word = ass.Word()
	} else {
		g.Word = wrd
	}

	g.Use(g.Word)

	return g
}

func GetGame(id string) *Game {
	return &Game{
		Id: id,
	}
}

func (g *Game) assoLog(from, to string) {
	log.Info("new asso: " + from + " → " + to + " ;")
}

func (g *Game) UserStep(word string, a string) bool {
	if g.Used(a) {
		return false
	}

	g.Use(a)

	asso.Add(word, a)

	return true
}

func (g *Game) ComputerStep(word string) *Record {
	list := ass.Asso(word)

	for _, v := range list {
		if !g.Used(v) {
			g.Use(v)

			res := &Record{
				Word: word,
				Asso: v,
			}

			res.Image, _ = avatars.Random()
			return res
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	words := sociation.GetWords(ctx, word)

	nwords := make([]string, len(words))
	copy(nwords, words)

	slice.Shuffle(nwords)

	for _, v := range nwords {

		w := text.Prepare(v)

		if !g.Used(w) {

			g.Use(w)

			r := &Record{
				Word: word,
				Asso: w,
			}

			r.Image, _ = avatars.Random()

			g.assoLog(word, w)

			return r
		}
	}

	return nil
}

func (g *Game) Used(word string) bool {
	return cache.HasWord(g.Id, word)
}

func (g *Game) Use(word string) {
	cache.SaveWord(g.Id, word)
}
