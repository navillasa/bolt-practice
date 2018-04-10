package boltdb

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/boltdb/bolt"
)

type File struct {
	EncryptedPath []string
	Data          []byte
	OldValHash    []byte
	Results       [][]string
	EndingPath    []string
	Truncated     bool
}

var (
	ErrFilenameTaken = errors.New("error filename taken")
)

func (client *Client) Create(file File) error {
	return client.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("files"))

		fileKey := []byte(file.EncryptedPath)

		v := b.Get(fileKey)
		if v != nil {
			return ErrFilenameTaken
		}

		fileBytes, err := json.Marshal(file)
		if err != nil {
			log.Println(err)
		}

		return b.Put(fileKey, fileBytes)
	})
}

func (client *Client) Get() {
}
