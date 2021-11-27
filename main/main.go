package main

import "dsl/def"

func main() {
	err := def.InitDef()
	if err != nil {
		return
	}
}
