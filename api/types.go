package api

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-02-08

type GameData struct {
	GameId string `json:"game_id"`
	Word   string `json:"word"`
	Before string `json:"word_before"`
}

type GameResponse struct {
	GameId  string `json:"game_id"`
	Word    string `json:"word"`
	Image   string `json:"image"`
	Victory bool   `json:"victory"`
}
