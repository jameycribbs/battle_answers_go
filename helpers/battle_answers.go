package helpers

import (
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

func PopulateBattleAnswerDisplays(db *mgo.Database, recs []models.BattleAnswerRec) []BattleAnswerDisplay {
	recsSize := len(recs)
	displays := make([]BattleAnswerDisplay, recsSize)

	for i, rec := range recs {
		displays[i] = populateBattleAnswerDisplay(db, rec)
	}

	return displays
}

func populateBattleAnswerDisplay(db *mgo.Database, rec models.BattleAnswerRec) BattleAnswerDisplay {
	var display BattleAnswerDisplay
	var game models.GameRec

	display.Question = rec.Question
	display.Answer = rec.Answer
	display.State = rec.State
	display.SubmitterEmail = rec.SubmitterEmail
	display.Tags = strings.Join(rec.Tags, " ")

	rec.FindGame(db, &game)

	display.GameName = game.Name

	return display
}
