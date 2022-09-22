package main

import (
	"fmt"
	"strings"

	xmlToJson "github.com/basgys/goxml2json"
	json "github.com/tidwall/gjson"
)

// process and return XML data
func getXml(data, query, defaultVal string) string {
	xml := strings.NewReader(data)
	jsonData, err := xmlToJson.Convert(xml)

	if err != nil {
		stop(fmt.Sprintf("error: XML conversion error. %s", err.Error()))
	}

	res := json.Get(jsonData.String(), query)

	if !res.Exists() {
		return defaultVal
	}

	return res.String()
}
