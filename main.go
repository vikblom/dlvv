package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

var (
	verbose = flag.Bool("v", false, "")
	count   = flag.Int("count", 0, "")
	run     = flag.String("run", "", "")
)

func main() {
	// TODO: Must fish out both dlv command and the debug target.
	cmd := os.Args[1]
	err := flag.CommandLine.Parse(os.Args[2:])
	if err != nil {
		panic(err)
	}
	for _, v := range flag.Args() {
		fmt.Println(":", v)
	}

	extra := flag.Args()
	if *verbose {
		extra = append(extra, "-test.v")
	}
	if *count > 0 {
		extra = append(extra, "-test.count")
		extra = append(extra, strconv.Itoa(*count))
	}
	if *run != "" {
		extra = append(extra, "-test.run")
		extra = append(extra, *run)
	}

	// Replace dlvv with dlv.
	dlv, err := exec.LookPath("dlv")
	if err != nil {
		panic(err)
	}
	env := os.Environ()

	args := append([]string{dlv, cmd}, flag.Args()...)
	if len(extra) > 0 {
		args = append(args, "--")
		args = append(args, extra...)
	}
	fmt.Println(args)
	err = syscall.Exec(dlv, args, env)
	if err != nil {
		panic(err)
	}
}
