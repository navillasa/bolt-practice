package main

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/navillasa/bolt-practice/storage/boltdb"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = logrus.StandardLogger()
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
		DisableSorting:  true,
	})
}

func main() {
	logger.Info("Serving on localhost:3000")

	bdb, err := boltdb.New()
	if err != nil {
		logger.Fatalln(boltdb.ErrInitDb, err)
		return
	}

	defer bdb.DB.Close()

	http.ListenAndServe(":3000", start())
}

func start() *httprouter.Router {
	router := httprouter.New()
	return router
}
