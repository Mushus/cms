package cms

type ContentStore interface {
	Save(id string, data ContentData) error
	Restore(id string) (ContentData, error)
	Close()
}

type ContentData interface {
	Encode([]byte) error
	Decode() ([]byte, error)
}
