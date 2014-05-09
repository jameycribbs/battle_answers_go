package controllers

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/jameycribbs/battle_answers/models"
	"labix.org/v2/mgo"
	"net/http"
)

type GameDisplay struct {
	Name string
}

type GameForm struct {
	Name string `form:"name"`
}

/////////////////////////////////////////////////////////////////////////////////////////////
// Controller Actions
/////////////////////////////////////////////////////////////////////////////////////////////
func GamesIndex(r render.Render, db *mgo.Database, req *http.Request) {
	var recs []models.GameRec

	recs = models.GetGameRecs(db, nil)

	templateData := map[string]interface{}{"metatitle": "Battle Answers", "currentPath": req.URL.Path,
		"recs": populateGameDisplays(db, recs)}
	r.HTML(200, "games/index", templateData)
}

func GamesNew(r render.Render, db *mgo.Database, req *http.Request) {
	templateData := map[string]interface{}{"metatitle": "Battle Answers", "currentPath": req.URL.Path}
	r.HTML(200, "games/new", templateData)
}

func GamesCreate(form GameForm, r render.Render, db *mgo.Database, req *http.Request) {
	var rec models.GameRec

	rec.Name = form.Name

	models.InsertGame(db, rec)

	GamesIndex(r, db, req)
}

/////////////////////////////////////////////////////////////////////////////////////////////
// Helper functions
/////////////////////////////////////////////////////////////////////////////////////////////
func populateGameDisplays(db *mgo.Database, recs []models.GameRec) []GameDisplay {
	recsSize := len(recs)
	displays := make([]GameDisplay, recsSize)

	for i, rec := range recs {
		displays[i] = populateGameDisplay(db, rec)
	}

	return displays
}

func populateGameDisplay(db *mgo.Database, rec models.GameRec) GameDisplay {
	var display GameDisplay

	display.Name = rec.Name

	return display
}
