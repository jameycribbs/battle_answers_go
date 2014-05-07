package models

import (
	"labix.org/v2/mgo"
)

type BattleAnswerRec struct {
	GameId         string
	Question       string
	Answer         string
	State          string
	SubmitterEmail string
	Tags           []string
}

type BattleAnswerForm struct {
	GameId         string `form:"gameid"`
	Question       string `form:"question"`
	Answer         string `form:"answer"`
	State          string `form:"state"`
	SubmitterEmail string `form:"submitterEmail"`
	Tags           string `form:"tags"`
}

func GetBattleAnswerRecs(db *mgo.Database, query interface{}) []BattleAnswerRec {
	var recs []BattleAnswerRec

	db.C("battle_answers").Find(query).All(&recs)

	return recs
}

/*
func insertBattleAnswer(db *mgo.Database, rec BattleAnswerRec) {
}
*/

func (rec BattleAnswerRec) FindGame(db *mgo.Database, game *GameRec) {
	FindGameById(db, rec.GameId, game)
}
