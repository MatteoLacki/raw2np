package main

import (
	"fmt"
	gonpy "github.com/kshedden/gonpy"
)

func main() {

	numSlice := []float64{5, 4, 3, 2, 1}

	fmt.Println(numSlice)
	w, _ := gonpy.NewFileWriter("some_data.npy")
	_ = w.WriteFloat64(numSlice)
	fmt.Println("Done")
}
