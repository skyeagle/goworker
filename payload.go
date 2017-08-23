package goworker

import "encoding/json"

type Payload struct {
	Class string          `json:"class"`
	Args  json.RawMessage `json:"args"`
}
