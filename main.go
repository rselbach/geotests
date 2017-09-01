package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var featureCollection *FeatureCollection

var (
	dataFile      string
	listen        string
	ignoreZeroPop bool // ignore features with 0 population
)

func init() {
	flag.StringVar(&dataFile, "filename", "canada_cities.geojson", "A geojson file containing the data")
	flag.StringVar(&dataFile, "f", "canada_cities.geojson", "A geojson file containing the data (shorthand)")
	flag.StringVar(&listen, "l", ":8000", "Where the server will listen to")
	flag.BoolVar(&ignoreZeroPop, "nz", false, "Ignore features with population 0")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	log.Printf("Reading data from %q\n", dataFile)
	var err error
	ignoreZeroPop = true
	featureCollection, err = NewFeatureCollectionFromFile("/home/robteix/go/src/github.com/robteix/geotests/data/canada_cities.geojson")
	//featureCollection, err = NewFeatureCollectionFromFile(dataFile)
	if err != nil {
		log.Fatalln("could create the feature collection: ", err)
	}

	log.Printf("Read and indexed %d features. All ready.\n", len(featureCollection.Features))
	log.Fatal(http.ListenAndServe(listen, setupAPIRouter()))
}

func usage() {
	fmt.Println("Usage: geotest [options]")
	flag.PrintDefaults()
}