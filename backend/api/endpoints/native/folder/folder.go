package folder

import (
	"os/exec"
)

func dir() (any, error) {
	cmd := exec.Command("cmd.exe", "/C", "dir")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return string(output), nil
}

func Choose() (any, error) {
	cmd := exec.Command("./bin/windows/winapi/winapi.exe")
	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	return string(output), nil
}
