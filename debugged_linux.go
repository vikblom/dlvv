package dlvv

import (
	"fmt"
	"os"
	"regexp"
	"syscall"
)

var re = regexp.MustCompile(`TracerPid:\s+([0-9]+)`)

func Debugged() (bool, error) {
	pid := syscall.Getpid()
	status, err := os.ReadFile(fmt.Sprintf("/proc/%d/status", pid))
	if err != nil {
		return false, fmt.Errorf("couldn't get proc tracer: %s", err)
	}
	sub := re.FindSubmatch(status)
	if sub == nil {
		return false, fmt.Errorf("couldn't find proc tracer PID")
	}
	if string(sub[1]) == "0" {
		return false, nil
	}
	return true, nil
}
