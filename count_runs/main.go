package main

import (
	"bitbucket.org/proteinspector/ms/unthermo"
	"flag"
	"fmt"
	"log"
)

func main() {
	var filename string
	var scans int
	//Parse arguments
	flag.StringVar(&filename, "raw", "small.RAW", "name of the RAW file")
	flag.Parse()

	//open RAW file
	file, err := unthermo.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scans = file.NScans()
	defer file.Close()

	//Print the Spectrum at the supplied scan number
	fmt.Println(scans)
}
