package main

import (
	//	"fmt"
	"github.com/codegangsta/martini"
	//"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/jameycribbs/battle_answers/controllers"
	"labix.org/v2/mgo"
)

func DB() martini.Handler {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	return func(c martini.Context) {
		s := session.Clone()
		c.Map(s.DB("battle_answers_db"))

		defer s.Close()
		c.Next()
	}
}

func main() {
	m := martini.Classic()

	m.Use(DB())

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout:    "layout",
	}))

	m.Get("/", func(r render.Render, db *mgo.Database) {
		templateData := map[string]interface{}{"metatitle": "Battle Answers"}
		r.HTML(200, "index", templateData)
	})

	/*
		m.Get("/games", func(r render.Render, db *mgo.Database) {
			templateData := map[string]interface{}{"metatitle": "Battle Answers", "games": GetAllGames(db)}
			r.HTML(200, "games/index", templateData)
		})

		m.Get("/games/new", func(r render.Render, db *mgo.Database) {
			templateData := map[string]interface{}{"metatitle": "Battle Answers"}
			r.HTML(200, "games/new", templateData)
		})

		m.Post("/games", binding.Form(Game{}), func(game Game, r render.Render, db *mgo.Database) {
			db.C("games").Insert(game)
			templateData := map[string]interface{}{"metatitle": "Battle Answers", "games": GetAllGames(db)}
			r.HTML(200, "games/index", templateData)
		})
	*/

	m.Get("/battle_answers", controllers.BattleAnswersIndex)

	/*
		m.Get("/battle_answers", func(r render.Render, db *mgo.Database) {
			templateData := map[string]interface{}{"metatitle": "Battle Answers", "battleanswers": GetBattleAnswerShows(db, nil)}
			r.HTML(200, "battle_answers/index", templateData)
		})

		m.Get("/battle_answers/new", func(r render.Render, db *mgo.Database) {
			templateData := map[string]interface{}{"metatitle": "Battle Answers", "games": GetAllGames(db)}
			r.HTML(200, "battle_answers/new", templateData)
		})

		m.Post("/battle_answers", binding.Form(BattleAnswer{}), func(battleAnswer BattleAnswer, r render.Render, db *mgo.Database) {
			battleAnswer.Tags = strings.Split(battleAnswer.Tags[0], " ")

			db.C("battle_answers").Insert(battleAnswer)
			templateData := map[string]interface{}{"metatitle": "Battle Answers", "battleanswers": GetBattleAnswerShows(db, nil)}
			r.HTML(200, "battle_answers/index", templateData)
		})
	*/

	m.Run()
}
