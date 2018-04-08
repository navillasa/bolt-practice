package boltdb

import (
	"errors"
	"github.com/boltdb/bolt"
	"time"
)

var (
	defaultTimeout = 1 * time.Second

	ErrDbOpen             = errors.New("error boltdb failed to open")
	ErrInitDb             = errors.New("error instantiating boltdb")
	ErrCreatingUserBucket = errors.New("error creating user bucket")
)

type Client struct {
	DB          *bolt.DB
	UsersBucket *bolt.Bucket
}

func New() (*Client, error) {
	db, err := bolt.Open("mystery.db", 0600, &bolt.Options{Timeout: defaultTimeout})
	if err != nil {
		return nil, ErrDbOpen
	}

	b := &bolt.Bucket{}
	err = db.Update(func(tx *bolt.Tx) error {
		b, err = tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return ErrCreatingUserBucket
		}
		return nil
	})

	return &Client{
		DB: db,
	}, nil
}
