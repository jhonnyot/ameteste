package model

import "github.com/globalsign/mgo/bson"

/*Planeta define a estrutura do objeto Planeta */
type Planeta struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Nome      string        `bson:"nome" json:"nome"`
	Clima     string        `bson:"clima" json:"clima"`
	Terreno   string        `bson:"terreno" json:"terreno"`
	Aparicoes int           `bson:"aparicoes" json:"aparicoes"`
}
