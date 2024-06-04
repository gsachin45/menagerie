package handlers

import (
	"context"
	"fmt"
	"menagerie/db"
	"menagerie/models"
	"menagerie/utils"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPetWithEventsHandler(w http.ResponseWriter, r *http.Request) {
	logger := utils.Logger
	vars := mux.Vars(r)
	petID := vars["id"]

	// Retrieve the pet
	var pet models.Pet
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id, _ := strconv.ParseInt(petID, 10, 64)

	err := db.PetsCollection.FindOne(ctx, bson.M{"id": int(id)}).Decode(&pet)
	if err != nil {
		logger.Error(err.Error())
		msg := fmt.Sprintf("unable to find pet with petId: %s", petID)
		respondWithError(w, http.StatusInternalServerError, msg)
		return
	}

	// Retrieve all events for the pet
	var events []models.Event
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := db.EventsCollection.Find(ctx, bson.M{"petid": petID})
	if err != nil {
		logger.Error(err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var event models.Event
		if err := cursor.Decode(&event); err != nil {
			logger.Error(err.Error())
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		events = append(events, event)
	}

	// Sort events by date in descending order (latest first)
	sort.SliceStable(events, func(i, j int) bool {
		return events[i].Date.After(events[j].Date)
	})

	// Combine pet and events into a single response
	petWithEvents := struct {
		Pet    models.Pet     `json:"pet"`
		Events []models.Event `json:"events"`
	}{
		Pet:    pet,
		Events: events,
	}

	respondWithJSON(w, http.StatusOK, petWithEvents)
}
