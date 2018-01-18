package cms

type ContentStore interface {
	Save(id string, data interface{}) error
	Restore(id string) (interface{}, error)
	Close()
}
