package handlers

import (
	"context"
	"encoding/json"
	"menagerie/db"
	"menagerie/models"
	"menagerie/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllPetsHandler(w http.ResponseWriter, r *http.Request) {
	var pets []models.Pet
	logger := utils.Logger

	species := r.URL.Query().Get("species")
	filter := bson.M{}
	if species != "" {
		filter["species"] = species
	}

	cursor, err := db.PetsCollection.Find(context.TODO(), filter)
	if err != nil {
		logger.Error(err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var pet models.Pet
		if err := cursor.Decode(&pet); err != nil {
			logger.Error(err.Error())
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		pets = append(pets, pet)
	}

	respondWithJSON(w, http.StatusOK, pets)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
