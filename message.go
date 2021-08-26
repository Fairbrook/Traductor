package main

import (
	"encoding/json"

	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

func handleMessages(w *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "process":
		var str string
		if len(m.Payload) > 0 {
			if err = json.Unmarshal(m.Payload, &str); err != nil {
				payload = err.Error()
				return
			}
		}
		if payload, err = processString(str); err != nil {
			payload = err.Error()
			return
		}
	}
	return
}
