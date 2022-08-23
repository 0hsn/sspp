package main

import (
	"bufio"
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
	var json, or, data string

	// define selector flag
	flag.StringVarP(&json, "json", "j", "", "valid data selector")

	// define default value flag
	flag.StringVar(&or, "or", "", "valid data selector")

	// define data flag
	flag.StringVar(&data, "data", "", "valid data selector")

	flag.Parse()

	cliopts.json = json
	cliopts.or = or

	if len(data) > 0 {
		cliopts.data = data
	} else if isInputFromPipe() {
		cliopts.data = readFromPipe()
	}
}

func convertOptionsToFeature(opts *options) *Feature {
	return &Feature{OpType: JSON, Query: opts.json, Data: opts.data, DefaulVal: opts.or}
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func readFromPipe() string {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	var buf []byte

	for scanner.Scan() {
		buf = append(buf, scanner.Bytes()...)
	}

	return string(buf[:])
}
