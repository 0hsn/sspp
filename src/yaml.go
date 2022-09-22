package main

import (
	"fmt"

	"github.com/ghodss/yaml"
	json "github.com/tidwall/gjson"
)

// process and return YAML data
func getYaml(data, query, defaultVal string) string {
	jsonData, err := yaml.YAMLToJSON([]byte(data))

	if err != nil {
		stop(fmt.Sprintf("error: YAML conversion error. %s", err.Error()))
	}

	res := json.Get(string(jsonData), query)

	if !res.Exists() {
		return defaultVal
	}

	return res.String()

}
