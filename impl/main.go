package main

import (
	"log"

	"github.com/Mushus/cms"
	"github.com/Mushus/cms/interface/web"
	"github.com/Mushus/cms/store"
	"github.com/boltdb/bolt"
)

func main() {
	cr := store.NewContentRegistry()
	cr.Register("markdown", Markdown{})

	db, err := bolt.Open("../content.db", 0666, nil)
	if err != nil {
		log.Fatalf("cannot open store: %v", err)
	}

	bstore, err := store.NewBoltStore(db)
	if err != nil {
		log.Fatalf("failed to open content store: %v", err)
	}
	bstore.SetContentRegistry(cr)

	c := cms.New()
	c.SetContentStore(bstore)

	s := web.NewServer(c)
	log.Fatal(s.Start())
}
