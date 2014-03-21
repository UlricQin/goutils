package systool

import (
	"os/exec"
	"bytes"
	"github.com/ulricqin/goutils/strtool"
)

func CmdOut(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func CmdOutNoLn(name string, arg ...string) (out string, err error) {
	out, err = CmdOut(name, arg...)
	if err != nil {
		return
	}

	return strtool.TrimRightSpace(string(out)), nil
}

