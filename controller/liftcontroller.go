package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/henryhiraki22/lifts/database"
	"github.com/henryhiraki22/lifts/entity"
)

func ListTrainingA(w http.ResponseWriter, r *http.Request) {
	liftsA := []entity.LiftA{}
	database.Connector.Find(&liftsA)
	json.NewEncoder(w).Encode(liftsA)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func CreateLiftA(w http.ResponseWriter, r *http.Request) {
	var liftA entity.LiftA
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &liftA)
	database.Connector.Create(&liftA)
	json.NewEncoder(w).Encode(liftA)
}

func ListTrainingB(w http.ResponseWriter, r *http.Request) {
	liftsB := []entity.LiftB{}
	database.Connector.Find(&liftsB)
	json.NewEncoder(w).Encode(liftsB)
}

func CreateLiftB(w http.ResponseWriter, r *http.Request) {
	var liftB entity.LiftB
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &liftB)
	database.Connector.Create(&liftB)
	json.NewEncoder(w).Encode(liftB)
}
