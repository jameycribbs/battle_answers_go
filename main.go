package main

import (
	//	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
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

	m.Get("/games", controllers.GamesIndex)
	m.Get("/games/new", controllers.GamesNew)
	m.Post("/games", binding.Form(controllers.GameForm{}), controllers.GamesCreate)

	m.Get("/battle_answers", controllers.BattleAnswersIndex)
	m.Get("/battle_answers/new", controllers.BattleAnswersNew)
	m.Post("/battle_answers", binding.Form(controllers.BattleAnswerForm{}), controllers.BattleAnswersCreate)

	m.Get("/search_answers", controllers.SearchAnswersIndex)
	m.Get("/search_answers/new", controllers.SearchAnswersNew)
	m.Post("/search_answers", binding.Form(controllers.SearchAnswerForm{}), controllers.SearchAnswersCreate)

	m.Run()
}
