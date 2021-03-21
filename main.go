package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

func help(w io.Writer) {
	fmt.Fprintf(w, "usage: %s <cmd> [args...]\n", os.Args[0])
	fmt.Fprint(w, `
commands:
	help              Print this message
	in    <duration>  Alert when the given duration has elapsed
	at    <time>      Alert at the given time
`)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
}

func run() error {
	args := os.Args[1:]
	if len(args) < 1 {
		help(os.Stderr)
		return nil
	}

	switch args[0] {
	case "-h", "--help", "help":
		help(os.Stdout)
		return nil
	case "in":
		if len(args) < 2 {
			help(os.Stderr)
			return errors.New("in: argument required")
		}

		if dur, err := time.ParseDuration(args[1]); err != nil {
			return fmt.Errorf("in: argument must be a valid duration: %v", err)
		} else {
			alertIn(dur)
		}
	case "at":
		if len(args) < 2 {
			help(os.Stderr)
			return errors.New("at: argument required")
		}

		if t, err := time.ParseInLocation("2006-01-02T15:04:05", args[1], time.Local); err != nil {
			return fmt.Errorf("at: argument must be a valid time: %v", err)
		} else {
			alertIn(t.Sub(time.Now()))
		}
	default:
		help(os.Stderr)
		return fmt.Errorf("invalid command: %s", args[0])
	}
	return nil
}

func alertIn(dur time.Duration) {
	fmt.Printf("Alerting in %v at %s\n", dur, time.Now().Add(dur).Local().Format("2006-01-02T15:04:05"))
	time.Sleep(dur)
	fmt.Println("Timer elapsed\u0007")
}
