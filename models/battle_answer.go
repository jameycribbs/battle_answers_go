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

func GetBattleAnswerRecs(db *mgo.Database, query interface{}) []BattleAnswerRec {
	var recs []BattleAnswerRec

	db.C("battle_answers").Find(query).All(&recs)

	return recs
}

func InsertBattleAnswer(db *mgo.Database, rec BattleAnswerRec) {
	db.C("battle_answers").Insert(rec)
}

func (rec BattleAnswerRec) FindGame(db *mgo.Database, game *GameRec) {
	FindGameById(db, rec.GameId, game)
}
