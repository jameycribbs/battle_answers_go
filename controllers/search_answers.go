package controllers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	"github.com/jameycribbs/battle_answers/helpers"
	"github.com/jameycribbs/battle_answers/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
	"strings"
)

type SearchAnswerForm struct {
	GameId   string `form:"gameid"`
	Keywords string `form:"keywords"`
}

/////////////////////////////////////////////////////////////////////////////////////////////
// Controller Actions
/////////////////////////////////////////////////////////////////////////////////////////////
func SearchAnswersIndex(r render.Render, db *mgo.Database, gameId string, keywords string, req *http.Request) {
	var game models.GameRec
	var recs []models.BattleAnswerRec
	var tags []string

	models.FindGameById(db, gameId, &game)

	tags = strings.Split(keywords, " ")

	recs = models.GetBattleAnswerRecs(db, bson.M{"gameid": gameId, "tags": bson.M{"$all": tags}})

	templateData := map[string]interface{}{"metatitle": "Battle Answers", "currentPath": req.URL.Path, "game": game,
		"keywords": keywords,
		"recs":     helpers.GetBattleAnswerDisplays(db, recs)}
	r.HTML(200, "search_answers/index", templateData)
}

func SearchAnswersNew(r render.Render, db *mgo.Database, req *http.Request, session sessions.Session) {
	lastGameIdSearched := session.Get("last_game_id_searched")

	templateData := map[string]interface{}{"metatitle": "Battle Answers", "currentPath": req.URL.Path,
		"lastGameIdSearched": lastGameIdSearched, "games": models.GetGameRecs(db, nil)}
	r.HTML(200, "search_answers/new", templateData)
}

func SearchAnswersCreate(form SearchAnswerForm, r render.Render, db *mgo.Database, req *http.Request, session sessions.Session) {
	session.Set("last_game_id_searched", form.GameId)

	SearchAnswersIndex(r, db, form.GameId, form.Keywords, req)
}
