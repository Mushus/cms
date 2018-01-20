package cms

type CMS interface {
	Edit(id string, data interface{})
	SetContentStore(cs ContentStore)
}
type cms struct {
	contentStore ContentStore
}

func New() CMS {
	return &cms{}
}

func (c *cms) SetContentStore(cs ContentStore) {
	c.contentStore = cs
}

func (c cms) Edit(id string, data interface{}) {

}
