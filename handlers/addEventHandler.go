package handlers

import (
	"context"
	"encoding/json"
	"menagerie/db"
	"menagerie/models"
	"menagerie/utils"
	"net/http"
)

func AddPetEventHandler(w http.ResponseWriter, r *http.Request) {
	logger := utils.Logger
	var event models.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := utils.Validate.Struct(event); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	_, err = db.EventsCollection.InsertOne(context.TODO(), event)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	logger.Info("added pet event in db ")
	respondWithJSON(w, http.StatusOK, nil)
}
