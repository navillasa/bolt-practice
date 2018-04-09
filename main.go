package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/navillasa/bolt-practice/storage/boltdb"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println("zap logger failed to instantiate")
		return
	}
	defer logger.Sync()
	logger.Info("serving on localhost:3000")

	bdb, err := boltdb.New()
	if err != nil {
		logger.Fatal("db error:",
			zap.Error(boltdb.ErrInitDb),
			zap.Error(err),
		)
		return
	}

	defer bdb.DB.Close()

	http.ListenAndServe(":3000", start())
}

func start() *httprouter.Router {
	router := httprouter.New()
	return router
}
