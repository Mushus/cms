package main

type Markdown struct {
	Text string
}

func (m *Markdown) Encode(data []byte) error {
	m.Text = string(data)
	return nil
}

func (m *Markdown) Decode() ([]byte, error) {
	return []byte(m.Text), nil
}
