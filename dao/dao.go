package dao

import (
	"log"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	model "github.com/jhonnyot/ameteste/model"
)

/*PlanetasDAO define a estrutura do DAO */
type PlanetasDAO struct {
	Servidor string
	Database string
}

/*Cria a variável banco*/
var banco *mgo.Database

/*Cria a constante Collection */
const (
	COLLECTION = "planetassw"
)

/*Connect conecta ao MongoDB */
func (dao *PlanetasDAO) Connect() {
	sessao, err := mgo.Dial(dao.Servidor)
	if err != nil {
		log.Fatal(err)
	}
	banco = sessao.DB(dao.Database)
}

/*Insert insere um novo planeta no banco */
func (dao *PlanetasDAO) Insert(planeta model.Planeta) error {
	err := banco.C(COLLECTION).Insert(&planeta)
	return err
}

/*FindAll lista todos os planetas */
func (dao *PlanetasDAO) FindAll() ([]model.Planeta, error) {
	var planetas []model.Planeta
	err := banco.C(COLLECTION).Find(bson.M{}).All(&planetas)
	return planetas, err
}

/*FindOneByName encontra um planeta pelo seu nome */
func (dao *PlanetasDAO) FindOneByName(Name string) (model.Planeta, error) {
	var planeta model.Planeta
	err := banco.C(COLLECTION).Find(bson.M{"nome": Name}).One(&planeta)
	return planeta, err
}

/*FindOneByID encontra um planeta dado seu ID */
func (dao *PlanetasDAO) FindOneByID(ID string) (model.Planeta, error) {
	var planeta model.Planeta
	err := banco.C(COLLECTION).FindId(bson.ObjectId(ID)).One(&planeta)
	return planeta, err
}

/*Delete exclui um planeta do banco */
func (dao *PlanetasDAO) Delete(planeta model.Planeta) error {
	err := banco.C(COLLECTION).Remove(&planeta)
	return err
}
