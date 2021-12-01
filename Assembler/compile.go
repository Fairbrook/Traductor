package Assembler

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/Fairbrook/analizador/Utils"
)

func (t *Translator) Compile(str string) (result string, err error) {
	errs := t.Translate(str)
	if len(errs) > 0 {
		err = errors.New(strings.Join(Utils.ErrorsToArray(errs), "\n"))
		return
	}
	splitName := strings.Split(t.Filename, ".")
	t.fileNameNoExt = t.Filename
	if len(splitName) > 0 {
		t.fileNameNoExt = splitName[0]
	}
	cmd := exec.Command("ml", "/c", "/Zc", "/coff", t.Filename)
	var stdout []byte
	stdout, err = cmd.Output()
	if err != nil {
		result = ""
		return
	}
	cmd = exec.Command("link", "/SUBSYSTEM:CONSOLE", t.fileNameNoExt+".obj")
	stdout, err = cmd.Output()
	if err != nil {
		result = ""
		return
	}
	result = string(stdout)
	return
}

func (t *Translator) CompileAndRun(str string) (result string, err error) {
	result, err = t.Compile(str)
	if err != nil {
		return
	}
	cmd := exec.Command(t.fileNameNoExt + ".exe")
	stdout, errors := cmd.Output()
	if errors != nil {
		result = ""
		return
	}
	result = string(stdout)
	return
}

func (t *Translator) TranslateAndOpen(str string) (result string, err error) {
	errs := t.Translate(str)
	if len(errs) > 0 {
		err = errors.New(strings.Join(Utils.ErrorsToArray(errs), "\n"))
		return
	}
	cmd := exec.Command("cmd", "/C", "start", t.Filename)
	stdout, err := cmd.Output()
	if err != nil {
		result = ""
		return
	}
	result = string(stdout)
	return
}
