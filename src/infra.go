package main

import (
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
	cliopts.data = data
}

func convertOptionsToFeature(opts *options) *Feature {
	return &Feature{OpType: JSON, Query: opts.json, Data: opts.data, DefaulVal: opts.or}
}
