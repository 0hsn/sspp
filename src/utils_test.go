package main

import (
	"testing"
)

func TestJsonDataBuilderGetIniPass(t *testing.T) {
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

func TestJsonDataBuilderGetIniFail(t *testing.T) {
	data := `
[small area]
var one = goesbyname
`

	ft := Feature{OpType: INI, Data: data, Query: "small area.var one", DefaultVal: ""}
	jdb := JsonDataBuilder{feature: &ft}

	if jdb.getIni() != "" {
		t.Error("Top level INI parse fail")
	}
}

func TestJsonDataBuilderGetJsonPass(t *testing.T) {
	data := `{
		"userId": 1,
		"id": 1,
		"title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		"body": {
			"asort": "gyp"
		}
	  }`
	ft := Feature{OpType: INI, Data: data, Query: "userId", DefaultVal: ""}
	jdb := JsonDataBuilder{feature: &ft}

	if jdb.getJson() != "1" {
		t.Error("Top level JSON parse fail")
	}
	jdb.feature.Query = "body.asort"
	if jdb.getJson() != "gyp" {
		t.Error("Subsequent level JSON parse fail")
	}
}

func TestJsonDataBuilderGetJsonFail(t *testing.T) {
	data := `"userId": 1,`
	ft := Feature{OpType: INI, Data: data, Query: "userId", DefaultVal: ""}
	jdb := JsonDataBuilder{feature: &ft}

	if jdb.getJson() != "Error: Invalid JSON." {
		t.Error("Top level JSON parse passed")
	}
}
