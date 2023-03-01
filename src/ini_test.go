package main

import (
	"testing"
)

func TestGetIni(t *testing.T) {
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

	if getIni(data, "title", "") != "translation_weblite_ca" {
		t.Error("Top level INI parse fail")
	}

	if getIni(data, "_database.name", "") != "mydb" {
		t.Error("Subsequent level INI parse fail")
	}
}
