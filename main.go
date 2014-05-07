package main

import (
	//	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strings"
)

type Game struct {
	Id   bson.ObjectId `bson:"_id,omitempty"`
	Name string        `form:"name"`
}

type BattleAnswer struct {
	GameId         string   `form:"gameid"`
	Question       string   `form:"question"`
	Answer         string   `form:"answer"`
	State          string   `form:"state"`
	SubmitterEmail string   `form:"submitterEmail"`
	Tags           []string `form:"tags"`
	VerifiedBy     string
	GameName       string `bson:"omitempty"`
}

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

func GetAllGames(db *mgo.Database) []Game {
	var games []Game
	db.C("games").Find(nil).All(&games)
	return games
}

func GetAllBattleAnswers(db *mgo.Database) []BattleAnswer {
	var battleAnswers []BattleAnswer
	var game Game

	db.C("battle_answers").Find(nil).All(&battleAnswers)

	// Populate GameName field in BattleAnswer struct from associated Game record.
	for i, b := range battleAnswers {
		db.C("games").FindId(bson.ObjectIdHex(b.GameId)).One(&game)
		b.GameName = game.Name
		battleAnswers[i] = b
	}

	return battleAnswers
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

	m.Get("/battle_answers", func(r render.Render, db *mgo.Database) {
		templateData := map[string]interface{}{"metatitle": "Battle Answers", "battleanswers": GetAllBattleAnswers(db)}
		r.HTML(200, "battle_answers/index", templateData)
	})

	m.Get("/battle_answers/new", func(r render.Render, db *mgo.Database) {
		templateData := map[string]interface{}{"metatitle": "Battle Answers", "games": GetAllGames(db)}
		r.HTML(200, "battle_answers/new", templateData)
	})

	m.Post("/battle_answers", binding.Form(BattleAnswer{}), func(battleAnswer BattleAnswer, r render.Render, db *mgo.Database) {
		battleAnswer.Tags = strings.Split(battleAnswer.Tags[0], " ")

		db.C("battle_answers").Insert(battleAnswer)
		templateData := map[string]interface{}{"metatitle": "Battle Answers", "battleanswers": GetAllBattleAnswers(db)}
		r.HTML(200, "battle_answers/index", templateData)
	})

	m.Get("/search_for_answers/new", func(r render.Render, db *mgo.Database) {
		templateData := map[string]interface{}{"metatitle": "Battle Answers", "games": GetAllGames(db)}
		r.HTML(200, "search_for_answers/new", templateData)
	})

	m.Run()
}
