package display

import "os/exec"

func TurnDisplayOff() error {
	return exec.Command("/bin/sh", "-c", "sleep 1; xset dpms force off").Run()
}
