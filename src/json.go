package main

import (
	json "github.com/tidwall/gjson"
)

// process and return JSON data
func getJson(data, query, defaultVal string) string {
	res := json.Get(data, query)

	if !res.Exists() {
		return defaultVal
	}

	return res.String()
}
