package main

import (
	"encoding/json"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/tidwall/gjson"
)

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
