package main

import "fmt"

func main() {
	feat := ParseFlags()
	jdb := JsonDataBuilder{feature: feat}
	fmt.Println(jdb.Export())
}
