package main

import (
	"encoding/json"
	"fmt"

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
		if payload, err = processString(str); err != nil {
			payload = err.Error()
			return
		}
	}
	// switch m.Name {
	// case "explore":
	// 	// Unmarshal payload
	// 	var path string
	// 	if len(m.Payload) > 0 {
	// 		// Unmarshal payload
	// 		if err = json.Unmarshal(m.Payload, &path); err != nil {
	// 			payload = err.Error()
	// 			return
	// 		}
	// 	}

	// 	// Explore
	// 	if payload, err = explore(path); err != nil {
	// 		payload = err.Error()
	// 		return
	// 	}
	// }
	return
}
