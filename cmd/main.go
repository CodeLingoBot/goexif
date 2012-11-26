
package main

import (
	"os"
	"fmt"
	"log"
	"flag"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

func main() {
	flag.Parse()
	fname := flag.Arg(0)

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	x.Walk(Walker{})
}

type Walker struct {}

func (_ Walker) Walk(name exif.FieldName, tag *tiff.Tag) error {
	data, _ := tag.MarshalJSON()
	fmt.Printf("%v: %v\n", name, string(data))
	return nil
}
