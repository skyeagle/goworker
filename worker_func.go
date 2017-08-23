package goworker

import "encoding/json"

type workerFunc func(string, *json.RawMessage) error
