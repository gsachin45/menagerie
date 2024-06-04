package handlers

import (
	"context"
	"encoding/json"
	"menagerie/db"
	"menagerie/models"
	"menagerie/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdatePetWithEventsHandler(w http.ResponseWriter, r *http.Request) {
	logger := utils.Logger
	vars := mux.Vars(r)
	petID := vars["id"]
	var pet models.Pet
	id, _ := strconv.ParseInt(petID, 10, 64)

	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := utils.Validate.Struct(pet); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	filter := bson.M{"id": id}

	updateResult := db.PetsCollection.FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": pet})
	if updateResult.Err() != nil {
		respondWithError(w, http.StatusInternalServerError, updateResult.Err().Error())
		return
	}
	pet.ID = int(id)
	logger.Info("updated pet data successfully")
	respondWithJSON(w, http.StatusOK, pet)
}
