package main

import (
	"encoding/json"
	"fmt"
	"strings"

	xmlToJson "github.com/basgys/goxml2json"

	"github.com/BurntSushi/toml"
	"github.com/ghodss/yaml"
	"github.com/tidwall/gjson"
	"github.com/wlevene/ini"
)

// process and return JSON data
func getIni(data, query, defaultVal string) string {
	iData := ini.New().Load([]byte(data))

	res := gjson.Get(string(iData.Marshal2Json()), query)

	if !res.Exists() {
		return defaultVal
	}

	return res.String()
}

// process and return JSON data
func getJson(data, query, defaultVal string) string {
	res := gjson.Get(data, query)

	if !res.Exists() {
		return defaultVal
	}

	return res.String()
}

func getToml(data, query, defaultVal string) string {

	var jsonI interface{}
	dataB := []byte(data)

	if err := toml.Unmarshal(dataB, &jsonI); err != nil {
		stop(fmt.Sprintf("error: TOML conversion error. %s", err.Error()))
	}

	jsonBytes, err := json.Marshal(interpolate(jsonI))
	if err != nil {
		stop(fmt.Sprintf("error: TOML conversion error [2]. %s", err.Error()))
	}

	jsonS := string(jsonBytes)
	res := gjson.Get(jsonS, query)

	if !res.Exists() {
		return defaultVal
	}

	return res.String()

}

// process and return XML data
func getXml(data, query, defaultVal string) string {
	xml := strings.NewReader(data)
	jsonData, err := xmlToJson.Convert(xml)

	if err != nil {
		stop(fmt.Sprintf("error: XML conversion error. %s", err.Error()))
	}

	res := gjson.Get(jsonData.String(), query)

	if !res.Exists() {
		return defaultVal
	}

	return res.String()
}

// process and return YAML data
func getYaml(data, query, defaultVal string) string {
	jsonData, err := yaml.YAMLToJSON([]byte(data))

	if err != nil {
		stop(fmt.Sprintf("error: YAML conversion error. %s", err.Error()))
	}

	res := gjson.Get(string(jsonData), query)

	if !res.Exists() {
		return defaultVal
	}

	return res.String()

}
