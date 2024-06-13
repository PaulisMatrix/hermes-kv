package main

import (
	"fmt"
	"hermeskv"
	"os"
)

func main() {
	s := hermeskv.GetNewKV(2)

	s.Set("hello", "world")
	s.Set("first", 100)
	s.Set("second", 200)

	val, err := s.Get("first")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("value received: ", val)
}
