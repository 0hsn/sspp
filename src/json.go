package main

import (
	json "github.com/tidwall/gjson"
)

func getJson(data, query, defaultVal string) string {
	res := json.Get(data, query)
	return res.String()
}
