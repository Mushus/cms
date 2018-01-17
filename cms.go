package cms

type cms struct {
	contentStore ContentStore
}

func New() *cms {
	return &cms{}
}

func (c *cms) SetContentStore(cs ContentStore) {
	c.contentStore = cs
}

func (c *cms) RegsterContent(typ string, value interface{}) {

}

func (c cms) Edit(id string, data interface{}) {

}
