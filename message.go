package main

import (
	"encoding/json"
	"fmt"

	"github.com/Fairbrook/analizador/Assembler"
	"github.com/Fairbrook/analizador/Semantico"
	"github.com/Fairbrook/analizador/Utils"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

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
			table, errs, _ := Semantico.Analize(str)
			if len(errs) > 0 {
				payload = Utils.ErrorsToArray(errs)
				return
			}
			payload = table.ToArray()
		}
	case "translate":
		{
			var str string
			if len(m.Payload) > 0 {
				if err = json.Unmarshal(m.Payload, &str); err != nil {
					payload = []string{err.Error()}
					return
				}
			}
			translator := Assembler.Translator{
				Filename: "output.asm",
			}
			translator.TranslateAndOpen(str)
		}
	case "compile":
		{
			var str string
			if len(m.Payload) > 0 {
				if err = json.Unmarshal(m.Payload, &str); err != nil {
					payload = []string{err.Error()}
					return
				}
			}
			translator := Assembler.Translator{
				Filename: "output.asm",
			}
			translator.Compile(str)
		}
	case "run":
		{
			var str string
			if len(m.Payload) > 0 {
				if err = json.Unmarshal(m.Payload, &str); err != nil {
					payload = []string{err.Error()}
					return
				}
			}
			translator := Assembler.Translator{
				Filename: "output.asm",
			}
			res, errs := translator.CompileAndRun(str)
			if errs != nil {
				payload = errs.Error()
				return
			}
			payload = res
			return
		}
	}
	return
}
