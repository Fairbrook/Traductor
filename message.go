package main

import (
	"encoding/json"
	"fmt"

	"github.com/Fairbrook/analizador/Semantico"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

func errorsToArray(errors []error) []string {
	res := []string{}
	for _, e := range errors {
		res = append(res, e.Error())
	}
	return res
}

func handleMessages(w *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	fmt.Printf(m.Name)
	switch m.Name {
	case "process":
		{
			var str string
			if len(m.Payload) > 0 {
				if err = json.Unmarshal(m.Payload, &str); err != nil {
					payload = []string{err.Error()}
					return
				}
			}
			fmt.Printf(str)
			table, errs := Semantico.Analize(str)
			if len(errs) > 0 {
				payload = errorsToArray(errs)
				return
			}
			payload = table.ToArray()
		}
	}
	return
}
