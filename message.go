package main

import (
	"encoding/json"
	"fmt"

	"github.com/Fairbrook/analizador/Semantico"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

func handleMessages(w *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	fmt.Printf(m.Name)
	switch m.Name {
	case "process":
		var str string
		if len(m.Payload) > 0 {
			if err = json.Unmarshal(m.Payload, &str); err != nil {
				payload = err.Error()
				return
			}
		}
		fmt.Printf(str)
		var errors []error
		if payload, errors = Semantico.ProcessString(str); len(errors) > 0 {
			payload = errors[0].Error()
			err = errors[0]
			return
		}
	}
	return
}
