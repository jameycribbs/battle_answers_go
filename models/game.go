package models

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type GameRec struct {
	Id   bson.ObjectId
	Name string
}

func GetGameRecs(db *mgo.Database, query interface{}) []GameRec {
	var recs []GameRec

	db.C("games").Find(query).All(&recs)

	return recs
}

func FindGameById(db *mgo.Database, id string, game *GameRec) {
	db.C("games").FindId(bson.ObjectIdHex(id)).One(&game)
}
