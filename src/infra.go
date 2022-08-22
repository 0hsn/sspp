package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

const (
	JSON = 0
)

type options struct {
	j, or, da string
}

type Feature struct {
	OpType          int8
	Data, DefaulVal string
}

func ParseFlags() *Feature {
	var cliopts = options{"", "", ""}
	defineFlags(&cliopts)
	fmt.Printf("%v\n", cliopts)

	return parseAndValidateFlags()
}

func defineFlags(cliopts *options) {
	var j, or, da string

	flag.StringVarP(&j, "json", "j", "", "valid data selector")
	flag.StringVar(&or, "or", "", "valid data selector")
	flag.StringVar(&da, "data", "", "valid data selector")

	flag.Parse()
	// fmt.Printf("%s\n", j)

	cliopts.j = j
	cliopts.or = or
	cliopts.da = da
}
func parseAndValidateFlags() *Feature {
	return &Feature{JSON, "", ""}
}
