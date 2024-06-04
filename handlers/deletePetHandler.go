package handlers

import (
	"context"
	"menagerie/db"
	"menagerie/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func DeletePetWithEventsHandler(w http.ResponseWriter, r *http.Request) {
	logger := utils.Logger
	vars := mux.Vars(r)
	petID := vars["id"]
	id, _ := strconv.ParseInt(petID, 10, 64)
	deleteResult, err := db.PetsCollection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	logger.Info("pet deleted successfully")
	respondWithJSON(w, http.StatusOK, deleteResult)
}
