package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/navillasa/bolt-practice/routes"
	"github.com/navillasa/bolt-practice/storage/boltdb"

	"go.uber.org/zap"
)

type Metadata interface {
	Get(encryptedPath []string) ([]byte, error)
	Put(encryptedPath []string, data []byte, oldValHash []byte) error
	List(startingPath, endingPath []string) (results [][]string, truncated bool, err error)
}

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

	file := routes.File{DB: bdb}

	http.ListenAndServe(":3000", start(file))
}

func start(f routes.File) *httprouter.Router {
	router := httprouter.New()

	router.GET("/file", f.Get)

	return router
}
