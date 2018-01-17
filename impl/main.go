package main

import (
	"log"

	"github.com/Mushus/cms"
	"github.com/Mushus/cms/store"
)

func main() {
	cr := store.ContentRegistory{
		"md": &Markdown{},
	}

	bstore, err := store.NewBoltStore("./content.db")
	if err != nil {
		log.Fatalf("failed to open content store: %v", err)
	}
	bstore.SetDataRegistory(cr)

	c := cms.New()
	c.SetContentStore(bstore)
	s := cms.NewServer(c)
	s.Start()
}
