package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"menagerie/db"
	"menagerie/models"
	"menagerie/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func AddPetsHandler(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	logger := utils.Logger
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := utils.Validate.Struct(pet); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	var data *models.Pet
	err = db.PetsCollection.FindOne(context.TODO(), bson.M{"id": pet.ID}).Decode(data)
	if err != nil || data != nil {
		msg := "pet already exists"
		logger.Error(msg)
		respondWithError(w, http.StatusConflict, msg)
		return
	}
	_, err = db.PetsCollection.InsertOne(context.TODO(), pet)
	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	logger.Infof("added pet %+v", pet)
	respondWithJSON(w, http.StatusOK, nil)
}
