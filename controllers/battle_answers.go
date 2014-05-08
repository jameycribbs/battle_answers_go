package controllers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/jameycribbs/battle_answers/helpers"
	"github.com/jameycribbs/battle_answers/models"
	"labix.org/v2/mgo"
	"strings"
)

type BattleAnswerForm struct {
	GameId         string `form:"gameid"`
	Question       string `form:"question"`
	Answer         string `form:"answer"`
	State          string `form:"state"`
	SubmitterEmail string `form:"submitterEmail"`
	Tags           string `form:"tags"`
}

/////////////////////////////////////////////////////////////////////////////////////////////
// Controller Actions
/////////////////////////////////////////////////////////////////////////////////////////////
func BattleAnswersIndex(r render.Render, db *mgo.Database) {
	var recs []models.BattleAnswerRec

	recs = models.GetBattleAnswerRecs(db, nil)

	templateData := map[string]interface{}{"metatitle": "Battle Answers", "recs": helpers.PopulateBattleAnswerDisplays(db, recs)}
	r.HTML(200, "battle_answers/index", templateData)
}

func BattleAnswersNew(r render.Render, db *mgo.Database) {
	templateData := map[string]interface{}{"metatitle": "Battle Answers", "games": models.GetGameRecs(db, nil)}
	r.HTML(200, "battle_answers/new", templateData)
}

func BattleAnswersCreate(form BattleAnswerForm, r render.Render, db *mgo.Database) {
	var rec models.BattleAnswerRec

	rec.GameId = form.GameId
	rec.Question = form.Question
	rec.Answer = form.Answer
	rec.State = form.State
	rec.SubmitterEmail = form.SubmitterEmail
	rec.Tags = strings.Split(form.Tags, " ")

	models.InsertBattleAnswer(db, rec)

	BattleAnswersIndex(r, db)
}
