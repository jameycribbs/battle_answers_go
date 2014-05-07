package controllers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/jameycribbs/battle_answers/models"
	"labix.org/v2/mgo"
	"strings"
)

type BattleAnswerDisplay struct {
	Question       string
	Answer         string
	State          string
	SubmitterEmail string
	Tags           string
	GameName       string
}

func BattleAnswersIndex(r render.Render, db *mgo.Database) {
	recs := models.GetBattleAnswerRecs(db, nil)

	templateData := map[string]interface{}{"metatitle": "Battle Answers", "recs": populateDisplays(db, recs)}
	r.HTML(200, "battle_answers/index", templateData)
}

/*
func NewBattleAnswer(r render.Render, db *mgo.Database) {
	templateData := map[string]interface{}{"metatitle": "Battle Answers", "games": GetAllGames(db)}
	r.HTML(200, "battle_answers/new", templateData)
}

func CreateBattleAnswer(battleAnswer BattleAnswer, r render.Render, db *mgo.Database) {
	battleAnswer.Tags = strings.Split(battleAnswer.Tags[0], " ")

	db.C("battle_answers").Insert(battleAnswer)
	templateData := map[string]interface{}{"metatitle": "Battle Answers", "battleanswers": GetBattleAnswerShows(db, nil)}
	r.HTML(200, "battle_answers/index", templateData)
}
*/

func populateDisplays(db *mgo.Database, recs []models.BattleAnswerRec) []BattleAnswerDisplay {
	recsSize := len(recs)
	displays := make([]BattleAnswerDisplay, recsSize)

	for i, rec := range recs {
		displays[i] = populateDisplay(db, rec)
	}

	return displays
}

func populateDisplay(db *mgo.Database, rec models.BattleAnswerRec) BattleAnswerDisplay {
	var display BattleAnswerDisplay
	var game models.GameRec

	display.Question = rec.Question
	display.Answer = rec.Answer
	display.State = rec.State
	display.SubmitterEmail = rec.SubmitterEmail
	display.Tags = strings.Join(rec.Tags, " ")

	rec.FindGame(db, &game)

	//	db.C("games").FindId(bson.ObjectIdHex(rec.GameId)).One(&game)

	display.GameName = game.Name

	return display
}
