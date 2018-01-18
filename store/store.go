package store

import "encoding/json"

type SaveData struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type RestoreData struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}
