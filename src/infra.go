package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	flag "github.com/spf13/pflag"
)

type options struct {
	json, or, data string
}

func ParseFlags() *Feature {
	var cliopts = options{"", "", ""}
	defineFlags(&cliopts)
	return convertOptionsToFeature(&cliopts)
}

func defineFlags(cliopts *options) {
	var json, xml, or, data string
	var hasJson, hasXml int8

	// define selector flag
	flag.StringVarP(&json, "json", "j", "", "valid dot-seperated data selector")
	flag.StringVarP(&xml, "xml", "x", "", "valid dot-seperated data selector")

	// define default value flag
	flag.StringVar(&or, "or", "", "valid data selector")

	// define data flag
	flag.StringVar(&data, "data", "", "valid data selector")

	flag.Parse()

	// validate data

	if json != "" {
		hasJson = 1
	}

	if xml != "" {
		hasXml = 1
	}

	if hasJson^hasXml == 0 {
		stop("error: Either multiple or no selector found")
	}

	// set data
	cliopts.json = json
	cliopts.or = or

	fi, err := os.Stdin.Stat()
	if err != nil {
		stop("error: While reading stdin")
	}

	if len(data) > 0 {
		cliopts.data = data
	} else if (fi.Mode() & os.ModeNamedPipe) != 0 {
		cliopts.data = readStdin()
	}

	if len(cliopts.data) == 0 {
		stop("error: No data found")
	}
}

func convertOptionsToFeature(opts *options) *Feature {
	return &Feature{OpType: JSON, Query: opts.json, Data: opts.data, DefaulVal: opts.or}
}

func readStdin() string {
	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	return string(output[:])
}

func stop(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(2)
}
