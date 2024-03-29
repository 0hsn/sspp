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
	var hasJson, hasXml, hasYaml, hasToml, hasIni int8
	var json, xml, yaml, toml, iniD string
	var or, data string

	// define selector flag
	flag.StringVarP(&json, "json", "J", "", "valid dot-separated json selector")
	flag.StringVarP(&xml, "xml", "X", "", "valid dot-separated xml selector")
	flag.StringVarP(&yaml, "yaml", "Y", "", "valid dot-separated yaml selector")
	flag.StringVarP(&toml, "toml", "T", "", "valid dot-separated toml selector")
	flag.StringVarP(&iniD, "ini", "I", "", "valid dot-separated ini selector")

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

	if toml != "" {
		hasToml = 1
	}

	if iniD != "" {
		hasIni = 1
	}

	if hasJson^hasXml^hasYaml^hasToml^hasIni == 0 {
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
	} else if hasToml == 1 {
		feature.Query = toml
		feature.OpType = TOML
	} else if hasIni == 1 {
		feature.Query = iniD
		feature.OpType = INI
	}

	// set default value
	feature.DefaultVal = or

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

// Read from stranded input stream
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

// Print and exit function: used for error
func stop(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(2)
}
