package main

import "fmt"

func main() {
	feat := ParseFlags()
	fmt.Println(GetDataFromFeature(feat))
}
