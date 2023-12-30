package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
	"syscall"
)

// splitTestArgs -v, -count and -run into extra.
// With the expected "-test.*" prefix.
func splitTestArgs(args []string) ([]string, []string) {
	extra := []string{}

	if i := slices.Index(args, "-v"); i >= 0 {
		extra = append(extra, "-test.v")
		args = append(args[:i], args[i+1:]...)
	}

	if i := slices.Index(args, "-count"); i >= 0 {
		extra = append(extra, "-test.count")
		extra = append(extra, args[i+1])
		args = append(args[:i], args[i+2:]...)
	}

	if i := slices.Index(args, "-run"); i >= 0 {
		extra = append(extra, "-test.run")
		extra = append(extra, args[i+1])
		args = append(args[:i], args[i+2:]...)
	}

	if i := slices.IndexFunc(args, func(s string) bool {
		return strings.HasPrefix(s, "-run=")
	}); i >= 0 {
		rest := strings.TrimPrefix(args[i], "-run=")
		extra = append(extra, "-test.run="+rest)
		args = append(args[:i], args[i+1:]...)
	}

	if i := slices.IndexFunc(args, func(s string) bool {
		return strings.HasPrefix(s, "-count=")
	}); i >= 0 {
		rest := strings.TrimPrefix(args[i], "-count=")
		extra = append(extra, "-test.count="+rest)
		args = append(args[:i], args[i+1:]...)
	}

	return args, extra
}

func main() {
	args, extra := splitTestArgs(os.Args)
	if len(extra) > 0 {
		args = append(args, "--")
		args = append(args, extra...)
	}

	// Replace dlvv with dlv.
	dlv, err := exec.LookPath("dlv")
	if err != nil {
		panic(err)
	}
	args[0] = dlv

	fmt.Println(args)
	err = syscall.Exec(dlv, args, os.Environ())
	if err != nil {
		panic(err)
	}
}
