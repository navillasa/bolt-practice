package routes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/navillasa/bolt-practice/storage/boltdb"
)

type File struct {
	DB *boltdb.Client
}

func (f *File) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "welcome to POST\n")

	file := boltdb.File{
		EncryptedPath: ps.ByName("name"),
		Data:          `file stuff`,
		OldValHash:    `here/i/am`,
		Results:       `here/i/am/at/this/here/path`,
		EndingPath:    `this/here/path`,
		Truncated:     false,
	}

	if err := f.DB.Create(file); err != nil {
		logger.Info("error:",
			zap.Error(err),
		)
	}
}

func (f *File) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "welcome to GET\n")

	f.DB.Get(ps.ByName("name"))

}
