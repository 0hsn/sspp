package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	flag "github.com/spf13/pflag"
)

// Create and load values to feature
func ParseFlags() *Feature {
	var feature = &Feature{}
	defineFlags(feature)
	return feature
}

// Define and process values form cli arguments
func defineFlags(feature *Feature) {
	var hasJson, hasXml, hasYaml int8
	var json, xml, yaml string
	var or, data string

	// define selector flag
	flag.StringVarP(&json, "json", "j", "", "valid dot-seperated json selector")
	flag.StringVarP(&xml, "xml", "x", "", "valid dot-seperated xml selector")
	flag.StringVarP(&yaml, "yaml", "y", "", "valid dot-seperated json selector")

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

	if yaml != "" {
		hasYaml = 1
	}

	if hasJson^hasXml^hasYaml == 0 {
		stop("error: Either multiple or no selector found")
	}

	// set selector
	if hasJson == 1 {
		feature.Query = json
		feature.OpType = JSON
	} else if hasXml == 1 {
		feature.Query = xml
		feature.OpType = XML
	} else if hasYaml == 1 {
		feature.Query = yaml
		feature.OpType = YAML
	}

	// set default value
	feature.DefaulVal = or

	// set data
	fi, err := os.Stdin.Stat()
	if err != nil {
		stop("error: While reading stdin")
	}

	if len(data) > 0 {
		feature.Data = data
	} else if (fi.Mode() & os.ModeNamedPipe) != 0 {
		feature.Data = readStdin()
	}

	if len(feature.Data) == 0 {
		stop("error: No data found")
	}
}

// Read from strandred input stream
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

// Print and exit funcion: use for error
func stop(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(2)
}
