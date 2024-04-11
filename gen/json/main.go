package main

import (
	"test/gen/json/enum_gen"
)

func main() {
	gen := enum_gen.EnumsGen()
	gen.Gen("test.json")
}
