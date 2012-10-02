package main

import (
	"fmt"
	"flag"
	"github.com/lukeroth/gdal"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	if filename == "" {
		fmt.Printf("Usage: test_tiff [filename]\n")
		return
	}
	fmt.Printf("Filename: %s\n", filename)
	
	fmt.Printf("Allocating buffer\n")
	var buffer [256 * 256]uint16
	//	buffer := make([]uint16, 256 * 256)
	
	fmt.Printf("Computing values\n")
	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			loc := x + y * 256
			val := x + y
			if val >= 256 {
				val -= 256
			}
			buffer[loc] = uint16(val)
		}
	}

	fmt.Printf("%d drivers available\n", gdal.GetDriverCount())

	fmt.Printf("Loading driver\n")
	driver, err := gdal.GetDriverByName("GTiff")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Creating dataset\n")
	dataset := driver.Create(filename, 256, 256, 1, gdal.UInt16, nil)
	defer dataset.Close()
	
	fmt.Printf("Getting raster band\n")
	raster := dataset.RasterBand(1)

	fmt.Printf("Writing to raster band\n")
	raster.IO(gdal.Write, 0, 0, 256, 256, buffer, 256, 256, 0, 0)
	
	fmt.Printf("End program\n")	
}