package store

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
)

type boltStore struct {
	db *bolt.DB
	cr *ContentRegistry
}

func NewBoltStore(db *bolt.DB) (*boltStore, error) {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("content"))
		return err
	})

	return &boltStore{
		db: db,
	}, nil
}

func (s *boltStore) SetContentRegistry(cr *ContentRegistry) {
	s.cr = cr
}

func (s boltStore) Save(id string, data interface{}) error {
	typ := s.cr.GetType(data)
	sv := SaveData{
		Type: typ,
		Data: data,
	}
	b, err := json.Marshal(sv)
	if err != nil {
		return fmt.Errorf("cannot encode save data: %v", err)
	}
	s.db.Update(func(tx *bolt.Tx) error {
		if content := tx.Bucket([]byte("content")); content != nil {
			if err := content.Put([]byte(id), b); err != nil {
				return err
			}
		}
		return nil
	})

	return nil
}

func (s boltStore) Restore(id string) (interface{}, error) {
	return nil, nil
}

func (s boltStore) Close() {
	s.db.Close()
}
