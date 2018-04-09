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

func (f *File) Get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "welcome to GET\n")
}
