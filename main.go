package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
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
			alertAt(time.Now().Add(dur))
		}
	case "at":
		if len(args) < 2 {
			help(os.Stderr)
			return errors.New("at: argument required")
		}

		if t, err := time.ParseInLocation("2006-01-02T15:04:05", args[1], time.Local); err != nil {
			return fmt.Errorf("at: argument must be a valid time: %v", err)
		} else {
			alertAt(t)
		}
	default:
		help(os.Stderr)
		return fmt.Errorf("invalid command: %s", args[0])
	}
	return nil
}

func alertAt(t time.Time) {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGABRT, syscall.SIGTERM)

	dur := t.Sub(time.Now())
	done := time.NewTimer(dur)
	tick := time.NewTicker(time.Second)

	fmt.Printf("Alerting at %s\n", time.Now().Add(dur).Local().Format("2006-01-02T15:04:05"))
	fmt.Print(dur.Round(time.Second))
	for {
		select {
		case <-done.C:
			clear()
			fmt.Println("Timer elapsed\u0007")
			return
		case <-tick.C:
			remaining := t.Sub(time.Now())
			clear()
			fmt.Print(remaining.Round(time.Second))
		case <-sigChan:
			fmt.Fprintln(os.Stderr, "\nAborted.")
			return
		}
	}
}

func clear() {
	// 60 width should be long enough to reliably catch any remnants
	// of previously printed durations. Without clearing we might
	// be stuck with a transition like '10s' => '9ss' => '8ss'
	// because the third column is never updated again.
	fmt.Printf("\r%60s\r", "")
}
