package main

import (
	json "github.com/tidwall/gjson"
	"github.com/wlevene/ini"
)

// process and return JSON data
func getIni(data, query, defaultVal string) string {
	iData := ini.New().Load([]byte(data))

	res := json.Get(string(iData.Marshal2Json()), query)

	if !res.Exists() {
		return defaultVal
	}

	return res.String()
}
