package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Set the log formatter
	log.SetFormatter(&log.TextFormatter{})
	log.WithFields(log.Fields{"time": time.Now().Format(time.RFC3339)}).
		Debug("Staring bashctl application")

	// Declare tall availabe flags (name, default, description)

	// Command flag
	var cmdstr string
	flag.StringVar(&cmdstr, "cmd", "", "Command to execute")
	flag.StringVar(&cmdstr, "c", "", "Command to execute (shorthand)")

	// Args flag
	var argstr string
	flag.StringVar(&argstr, "args", "", "Arguments to execute with the command")
	flag.StringVar(&argstr, "a", "", "Arguments to execute with the command (shorthand)")

	// Visibility flag
	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "View all the logs and traces")
	flag.BoolVar(&verbose, "v", false, "View all the logs and traces (shorthand)")

	// Parse all the flags extracted from command-line
	flag.Parse()

	// Set the level of detail depending on the verbose flags
	if verbose {
		log.SetLevel(logrus.DebugLevel)
	}

	// Log current args used (by command-line or used by default)
	log.WithFields(log.Fields{"cmd": cmdstr, "args": argstr}).Debug("Command line Params")
	log.WithFields(log.Fields{"args": flag.Args()}).Debug("Command line Args")

	// Check minimun flags have been injected
	if cmdstr == "" {
		log.WithFields(log.Fields{"time": time.Now().Format(time.RFC3339)}).
			Debug("Error: it must be set the command executable to run")
		fmt.Println("Error: it must be set the command executable to run")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Execute the command
	args := strings.Split(argstr, ",")
	cmd := exec.Command(cmdstr, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(out))

	log.WithFields(log.Fields{"time": time.Now().Format(time.RFC3339)}).Debug("Finish bashctl application")
}
