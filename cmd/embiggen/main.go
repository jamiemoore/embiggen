package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/alecthomas/units"
)

func main() {
	args := os.Args
	infile := args[1]
	outfile := args[2]
	base := args[3]

	var err error
	var f *os.File
	var targetFileSize int64

	thisthen, err := units.ParseBase2Bytes(base)
	targetFileSize = int64(thisthen)
	log.Println(int(thisthen), "Target file size is ", base, " which is ", targetFileSize, " bytes")

	f, err = os.Create(outfile)
	if err != nil {
		log.Fatal(err)
	}

	// If the file doesn't exist, create it, or append to the file
	if f, err = os.OpenFile(outfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal("error", err)
		}
	}()

	var current *os.File
	if current, err = os.OpenFile(infile, os.O_RDONLY, 0); err != nil {
		log.Fatal("error", err)
	}

	defer func() {
		if err := current.Close(); err != nil {
			log.Fatal("error", err)
		}
	}()

	if fileBytes, err := ioutil.ReadAll(current); err != nil {
		log.Fatal("error", err)
	} else {

		var fileSize int64
		for fileSize < targetFileSize {
			if _, err := f.Write([]byte(fileBytes)); err != nil {
				log.Fatal(err)
			}
			fi, err := f.Stat()
			if err != nil {
				log.Fatal(err)
			}
			fileSize = fi.Size()
			log.Print("filesize:", fileSize)
		}
	}
	log.Println("done")
}
