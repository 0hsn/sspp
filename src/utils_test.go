package main

import (
	"testing"
)

func TestJsonDataBuilderGetIni(t *testing.T) {
	data := `
;;Configuration settings for application
title=translation_weblite_ca
scriptUrl=http://translation.weblite.ca/index.php
multilingual_content=1

[_database]
	host=localhost
	name=mydb
	user=mydbuser
	password=foo	
`
	ft := Feature{OpType: INI, Data: data, Query: "title", DefaultVal: ""}
	jdb := JsonDataBuilder{feature: &ft}

	if jdb.getIni() != "translation_weblite_ca" {
		t.Error("Top level INI parse fail")
	}
	jdb.feature.Query = "_database.name"
	if jdb.getIni() != "mydb" {
		t.Error("Subsequent level INI parse fail")
	}
}
