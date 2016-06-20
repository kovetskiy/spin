package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/kovetskiy/godocs"
	"github.com/kovetskiy/spinner-go"
)

var (
	version = `1`
	usage   = `spin ` + version + `

Usage:
	spin [options]

Options:
    -i --stdin-as-indicator   Use stdin data as progress indicator.
    -o --write-stdin          Write stdin to spinner's stdout on exit.
    -s --status <string>      Use specified <string> as spinner status.
    -t --interval <millisec>  Use specified <millisec> as spinner iteration interval.
                               [default: 100]
    -h --help                 Show this screen.
    --version                 Show version.
`
)

func main() {
	args := godocs.MustParse(usage, version, godocs.UsePager)

	var (
		watchStdin  = args["--stdin-as-indicator"].(bool)
		writeStdin  = args["--write-stdin"].(bool)
		status, _   = args["--status"].(string)
		interval, _ = strconv.Atoi(args["--interval"].(string))

		ticker = time.NewTicker(
			time.Duration(int64(interval)) * time.Millisecond,
		)
		stdout = bytes.NewBuffer(nil)
	)

	spinner.SetStatus(status).SetOutput(os.Stdout).SetActive(true)

	if status != "" {
		fmt.Print(status)
	}

	defer func() {
		spinner.SetActive(false)
		spinner.Spin()

		if writeStdin {
			fmt.Print(stdout.String())
		}
	}()

	stdin, errors := getStdin()

	var ticking bool
	var reading bool

	for {
		select {
		case <-ticker.C:
			ticking = true

		case buffer := <-stdin:
			if writeStdin {
				stdout.WriteString(buffer)
			}
			reading = true

		case err := <-errors:
			if err == io.EOF {
				return
			}

			log.Println(err)
		}

		if ticking && ((watchStdin && reading) || (!watchStdin)) {
			spinner.Spin()

			ticking = false
			reading = false
		}
	}

}

func getStdin() (chan string, chan error) {
	var (
		reader  = make(chan string)
		errorer = make(chan error)
	)

	go func() {
		for {
			buffer := make([]byte, 0xFFFF)
			_, err := os.Stdin.Read(buffer)
			if err != nil {
				errorer <- err
				return
			}

			reader <- string(buffer)
		}
	}()

	return reader, errorer
}
