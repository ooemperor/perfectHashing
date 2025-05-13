package main

import (
	"fmt"
	"perfectHashing/src"
)

func main() {
	fmt.Println("Starting the perfect hashing main method.")

	hasher := src.MinimalPerfectHash{}

	e1 := src.ResultEntry{Value: "v1"}
	e2 := src.ResultEntry{Value: "v2"}
	e3 := src.ResultEntry{Value: "v4"}
	e4 := src.ResultEntry{Value: "v1"}
	e5 := src.ResultEntry{Value: "v2"}

	r1 := src.ResultSet{Entries: []*src.ResultEntry{&e1, &e2, &e3}}
	r2 := src.ResultSet{Entries: []*src.ResultEntry{&e4, &e5}}

	hasher.Build(&r1, &r2)

}
