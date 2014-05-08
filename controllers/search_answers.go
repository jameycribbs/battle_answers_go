package controllers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/jameycribbs/battle_answers/helpers"
	"github.com/jameycribbs/battle_answers/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strings"
)

type SearchAnswerForm struct {
	GameId   string `form:"gameid"`
	Keywords string `form:"keywords"`
}

/////////////////////////////////////////////////////////////////////////////////////////////
// Controller Actions
/////////////////////////////////////////////////////////////////////////////////////////////
func SearchAnswersIndex(r render.Render, db *mgo.Database, gameId string, keywords []string) {
	var recs []models.BattleAnswerRec

	recs = models.GetBattleAnswerRecs(db, bson.M{"gameid": gameId, "tags": bson.M{"$all": keywords}})

	templateData := map[string]interface{}{"metatitle": "Battle Answers", "recs": helpers.PopulateBattleAnswerDisplays(db, recs)}
	r.HTML(200, "search_answers/index", templateData)
}

func SearchAnswersNew(r render.Render, db *mgo.Database) {
	templateData := map[string]interface{}{"metatitle": "Battle Answers", "games": models.GetGameRecs(db, nil)}
	r.HTML(200, "search_answers/new", templateData)
}

func SearchAnswersCreate(form SearchAnswerForm, r render.Render, db *mgo.Database) {
	var keywords []string

	keywords = strings.Split(form.Keywords, " ")

	SearchAnswersIndex(r, db, form.GameId, keywords)
}