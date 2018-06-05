package main

import (
	"bitbucket.org/proteinspector/ms"
	"bitbucket.org/proteinspector/ms/unthermo"
	"flag"
	"fmt"
	gonpy "github.com/kshedden/gonpy"
	"log"
)

func main() {
	var scannumber int
	var filename string

	//Parse arguments
	flag.IntVar(&scannumber, "sn", 1, "the scan number")
	flag.StringVar(&filename, "raw", "small.RAW", "name of the RAW file")
	flag.Parse()

	//open RAW file
	file, err := unthermo.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Print the Spectrum at the supplied scan number
	scan2numpy(file.Scan(scannumber), scannumber)
}

//Save m/z and intensities to a numpy array.
func scan2numpy(scan ms.Scan, scannumber int) {
	mz := make([]float64, 0)
	in := make([]float64, 0)

	for _, peak := range scan.Spectrum() {
		mz = append(mz, peak.Mz)
		in = append(in, float64(peak.I))
	}

	table2npy(mz, fmt.Sprintf("%s%d%s", "mz_scan_", scannumber, ".npy"))
	table2npy(in, fmt.Sprintf("%s%d%s", "intensity_scan_", scannumber, ".npy"))

	fmt.Println("Writing to Numpy complete.")
}

// write a numpy file
func table2npy(numbers []float64, path string) {
	w, _ := gonpy.NewFileWriter(path)
	_ = w.WriteFloat64(numbers)
}
