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

type JsonDataBuilder struct {
	// Builds json data out of different data type
	feature *Feature
}

func (jdb *JsonDataBuilder) Export() string {
	switch jdb.feature.OpType {
	case JSON:
		return jdb.getJson()
	case XML:
		return jdb.getXml()
	case YAML:
		return jdb.getYaml()
	case TOML:
		return jdb.getToml()
	case INI:
		return jdb.getIni()
	default:
		return ""
	}
}

// process and return JSON data
func (jdb *JsonDataBuilder) getIni() string {
	iData := ini.New().Load([]byte(jdb.feature.Data))

	res := gjson.GetBytes(iData.Marshal2Json(), jdb.feature.Query)

	if !res.Exists() {
		return jdb.feature.DefaultVal
	}

	return res.String()
}

// process and return JSON data
func (jdb *JsonDataBuilder) getJson() string {
	if !gjson.Valid(jdb.feature.Data) {
		return "Error: Invalid JSON."
	}

	res := gjson.Get(jdb.feature.Data, jdb.feature.Query)

	if !res.Exists() {
		return jdb.feature.DefaultVal
	}

	return res.String()
}

func (jdb *JsonDataBuilder) getToml() string {

	var jsonI interface{}
	dataB := []byte(jdb.feature.Data)

	if err := toml.Unmarshal(dataB, &jsonI); err != nil {
		return fmt.Sprintf("error: TOML conversion error. %s", err.Error())
	}

	jsonBytes, err := json.Marshal(interpolate(jsonI))
	if err != nil {
		return fmt.Sprintf("error: TOML conversion error [2]. %s", err.Error())
	}

	jsonS := string(jsonBytes)
	res := gjson.Get(jsonS, jdb.feature.Query)

	if !res.Exists() {
		return jdb.feature.DefaultVal
	}

	return res.String()
}

// process and return XML data
func (jdb *JsonDataBuilder) getXml() string {
	xml := strings.NewReader(jdb.feature.Data)
	jsonData, err := xmlToJson.Convert(xml)

	if err != nil {
		return fmt.Sprintf("error: XML conversion error. %s", err.Error())
	}

	res := gjson.Get(jsonData.String(), jdb.feature.Query)

	if !res.Exists() {
		return jdb.feature.DefaultVal
	}

	return res.String()
}

// process and return YAML data
func (jdb *JsonDataBuilder) getYaml() string {
	jsonData, err := yaml.YAMLToJSON([]byte(jdb.feature.Data))

	if err != nil {
		return fmt.Sprintf("error: YAML conversion error. %s", err.Error())
	}

	res := gjson.Get(string(jsonData), jdb.feature.Query)

	if !res.Exists() {
		return jdb.feature.DefaultVal
	}

	return res.String()
}
