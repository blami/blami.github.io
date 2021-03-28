package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/blami/blami.github.io/api/db"
	"github.com/blami/blami.github.io/api/log"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var xdb db.DB

func main() {
	// Whether or not we run in GAE environment
	version, gae := os.LookupEnv("GAE_VERSION")
	if version == "" {
		version = "development"
	}

	// Setup logging (force JSON logging when in GAE)
	log.Setup(gae)
	log.WithFields(logrus.Fields{
		"version": version,
	}).Info("apisvc is starting")

	// Load configuration
	conf := Conf{
		DB:    "datastore",
		Port:  80,
		Debug: true,
		GAE:   gae,
	}
	conf.Load()
	if conf.Debug {
		log.SetLevel(logrus.DebugLevel)
	}

	// Setup database
	var err error
	xdb, err = db.MakeDB(conf.DB)
	if err != nil {
		log.WithError(err).Fatal("unable to setup db, exiting!")
	}

	r := mux.NewRouter()

	// Setup routes
	//r.HandleFunc("/")
	r.HandleFunc("/status", StatusHandler).Methods("GET")
	r.HandleFunc("/test/{key}/{value}", WHandler)
	r.HandleFunc("/test/{key}", RHandler)

	// Run server
	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%d", conf.Port),
	}
	s.ListenAndServe()
}

func WHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)

	if err := xdb.Set(ctx, vars["key"], vars["value"]); err != nil {
		log.WithError(err).WithFields(logrus.Fields{
			"key":   vars["key"],
			"value": vars["value"],
		}).Error("unable to store to db")
	}
}

func RHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)

	value, err := xdb.Get(ctx, vars["key"])
	if err != nil {
		log.WithError(err).WithFields(logrus.Fields{
			"key": vars["key"],
		}).Error("unable to read from db")
	}

	out := struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{
		vars["key"],
		value,
	}
	body, _ := json.Marshal(out)
	w.Write(body)
}

// Status handler
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	log.Info("request served")

	status := struct {
		Alive bool   `json:"alive"`
		Test  string `json:"test"`
	}{
		true,
		os.Getenv("GAE_VERSION"),
	}
	body, _ := json.Marshal(status)
	w.Write(body)
}
