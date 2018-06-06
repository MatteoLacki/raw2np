package main

import (
	"bitbucket.org/proteinspector/ms/unthermo"
	"flag"
	"fmt"
	gonpy "github.com/kshedden/gonpy"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//Parse arguments
	var path string
	// pass a pointer to "path" here: target gets a new value
	flag.StringVar(&path, "path", "", "path to where to save things.")
	flag.Parse()
	filenames := flag.Args()

	//save scans
	for _, filename := range filenames {
		scans2numpy(filename, path)
	}
}

// write a numpy file
func table2npy(numbers []float64, path string) {
	w, _ := gonpy.NewFileWriter(path)
	_ = w.WriteFloat64(numbers)
}

// save mz and intensities
func save(mz []float64, in []float64, path string) {
	os.MkdirAll(path, os.ModePerm)
	table2npy(mz, filepath.Join(path, "mz.npy"))
	table2npy(in, filepath.Join(path, "in.npy"))
}

// save scan individually and as aggregates
func scans2numpy(filename string, path string) {
	//open RAW file
	file, err := unthermo.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // defers file.Close until scans2numpy returns

	// parse filename
	path = filepath.Join(path,
		filename[0:len(filename)-len(filepath.Ext(filename))])

	// results
	mz_agg := make([]float64, 0)
	in_agg := make([]float64, 0)
	mz_scan := make([]float64, 0)
	in_scan := make([]float64, 0)

	for i := 1; i <= file.NScans(); i++ {
		scan := file.Scan(i)
		for _, peak := range scan.Spectrum() {
			mz_scan = append(mz_scan, peak.Mz)
			in_scan = append(in_scan, float64(peak.I))
		}

		// write scan data
		save(mz_scan, in_scan, filepath.Join(path, fmt.Sprintf("%d", i)))

		// updating the aggregated value
		mz_agg = append(mz_agg, mz_scan...)
		in_agg = append(in_agg, in_scan...)

		// reseting scan containers
		mz_scan = nil
		in_scan = nil
	}

	save(mz_agg, in_agg, path)
	fmt.Println("completed conversion to", path)
}
