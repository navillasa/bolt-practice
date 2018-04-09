package boltdb

type File struct {
	encryptedPath []string
	data          []byte
	oldValHash    []byte
	results       [][]string
	truncated     bool
}

func (bdb *Client) Get() {
}
