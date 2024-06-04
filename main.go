package main

import (
	"menagerie/db"
	"menagerie/handlers"
	"menagerie/utils"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

var log = logrus.New()

func main() {
	utils.Logger = log
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	router.HandleFunc("/pets", handlers.GetAllPetsHandler).Methods("GET")
	router.HandleFunc("/pets/{id}", handlers.GetPetWithEventsHandler).Methods("GET")
	router.HandleFunc("/pets", handlers.AddPetsHandler).Methods("POST")
	router.HandleFunc("/pets/{id}", handlers.UpdatePetWithEventsHandler).Methods("PUT")
	router.HandleFunc("/pets/{id}", handlers.AddPetEventHandler).Methods("POST")
	router.HandleFunc("/pets/{id}", handlers.UpdatePetWithEventsHandler).Methods("DELETE")
	db.ConnectToDB() //create connection with database
	log.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request
		log.Infof("[%s] %s %s %s", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path, r.RemoteAddr)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
