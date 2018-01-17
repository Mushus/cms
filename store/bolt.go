package store

import (
	"fmt"

	"github.com/Mushus/cms"
	"github.com/boltdb/bolt"
)

type boltStore struct {
	db *bolt.DB
	cr ContentRegistory
}

func NewBoltStore(path string) (*boltStore, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot open store: %v", err)
	}

	return &boltStore{
		db: db,
	}, nil
}

func (s *boltStore) SetDataRegistory(cr ContentRegistory) {
	s.cr = cr
}

func (s boltStore) Save(id string, data cms.ContentData) error {
	return nil
}

func (s boltStore) Restore(id string) (cms.ContentData, error) {
	return nil, nil
}

func (s boltStore) Close() {
	s.db.Close()
}
