package main

import (
	"bitbucket.org/proteinspector/ms/unthermo"
	"flag"
	"fmt"
	gonpy "github.com/kshedden/gonpy"
	"log"
	"path/filepath"
)

func main() {
	var filename string

	//Parse arguments
	flag.StringVar(&filename, "raw", "small.RAW", "name of the RAW file")
	flag.Parse()

	//save scans
	scans2numpy(filename)
}

// write a numpy file
func table2npy(numbers []float64, path string) {
	w, _ := gonpy.NewFileWriter(path)
	_ = w.WriteFloat64(numbers)
}

// save m/z
func scans2numpy(filename string) {
	//open RAW file
	file, err := unthermo.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// results
	mz := make([]float64, 0)
	in := make([]float64, 0)

	for i := 1; i <= file.NScans(); i++ {
		scan := file.Scan(i)
		for _, peak := range scan.Spectrum() {
			mz = append(mz, peak.Mz)
			in = append(in, float64(peak.I))
		}
	}

	extension := filepath.Ext(filename)
	name := filename[0 : len(filename)-len(extension)]

	table2npy(mz, fmt.Sprintf("%s%s%s", "mz_", name, ".npy"))
	table2npy(in, fmt.Sprintf("%s%s%s", "intensity_", name, ".npy"))

	fmt.Println("Writing to Numpy complete.")
}
