package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"

	cfg "github.com/jhonnyot/ameteste/config"

	DAO "github.com/jhonnyot/ameteste/dao"
	model "github.com/jhonnyot/ameteste/model"
)

var dao = DAO.PlanetasDAO{}
var config = cfg.Config{}

func init() {
	config.Read()

	dao.Servidor = config.Servidor
	dao.Database = config.Database
	dao.Connect()
}

/*respondWithError é o handler de respostas de erro */
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

/*respondWithJSON é o handler padrão de respostas */
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

/*AddPlanetEndPoint Adiciona um novo planeta */
func AddPlanetEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var planeta model.Planeta
	if err := json.NewDecoder(r.Body).Decode(&planeta); err != nil {
		respondWithError(w, http.StatusBadRequest, "Payload inválido da requisição. Verifique o objeto JSON da chamada.")
		return
	}
	planeta.ID = bson.NewObjectId()
	if err := dao.Insert(planeta); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, planeta)
}

/*FindAllEndPoint Lista todos os planetas */
func FindAllEndPoint(w http.ResponseWriter, r *http.Request) {
	planetas, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, planetas)
}

/*FindOneByNameEndPoint Encontra um planeta por nome */
func FindOneByNameEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	planeta, err := dao.FindOneByName(params["nome"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, planeta)
}

/*FindOneByIDEndPoint Encontra um planeta por ID */
func FindOneByIDEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	planeta, err := dao.FindOneByName(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, planeta)
}

/*DeletePlanetEndPoint Deleta um planeta */
func DeletePlanetEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var planeta model.Planeta
	if err := json.NewDecoder(r.Body).Decode(&planeta); err != nil {
		respondWithError(w, http.StatusBadRequest, "Payload inválido da requisição. Verifique o objeto JSON da chamada.")
		return
	}
	if err := dao.Delete(planeta); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/planetas", AddPlanetEndPoint).Methods("POST")
	router.HandleFunc("/planetas", FindAllEndPoint).Methods("GET")
	router.HandleFunc("/planetas/{nome}", FindOneByNameEndPoint).Methods("GET")
	router.HandleFunc("/planetas/{id}", FindOneByIDEndPoint).Methods("GET")
	router.HandleFunc("/planetas", DeletePlanetEndPoint).Methods("DELETE")

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
