package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	"github.com/jameycribbs/battle_answers/controllers"
	"html/template"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
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

	store := sessions.NewCookieStore([]byte("battle_answers"))
	m.Use(sessions.Sessions("my_battle_answers_session", store))

	m.Use(DB())

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout:    "layout",
		Funcs: []template.FuncMap{
			{
				"addInClass": func(args ...interface{}) string {
					var i int
					var className string

					i = args[0].(int)

					if i == 0 {
						className = " in"
					} else {
						className = ""
					}

					return className
				},
				"addSelected": func(args ...interface{}) string {
					var lastGameIdSearched string
					selectedString := ""

					lastGameIdSearched = args[1].(string)

					if args[0] == bson.ObjectIdHex(lastGameIdSearched) {
						selectedString = "selected"
					}

					return selectedString
				},
				"addActiveClass": func(args ...interface{}) string {
					className := ""

					for _, path := range args[1:] {
						if path == args[0] {
							className = "active"
							break
						}
					}

					return className
				},
			},
		},
	}))

	m.Get("/", func(r render.Render, db *mgo.Database, req *http.Request) {
		templateData := map[string]interface{}{"metatitle": "Battle Answers", "currentPath": req.URL.Path}
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
