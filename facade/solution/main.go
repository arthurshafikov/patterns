package main

import (
	"fmt"

	"github.com/arthurshafikov/patterns/facade/solution/facade"
)

func main() {
	// these 2 examples might be anywhere in your program
	exampleConvertMP4ToOgg()
	exampleConvertOggToMP4()
	/*
		File was converted successfully! Output - some-file.ogg
		File was converted successfully! Output - some-file.mp4
	*/
}

func exampleConvertMP4ToOgg() {
	// we don't know anything about the external library
	filename := "some-file.mp4"
	outputFilename := "some-file.ogg"
	wantFormat := "ogg"

	result := facade.NewVideoConverter().Convert(filename, outputFilename, wantFormat)
	if !result {
		panic("file could not be saved")
	}

	fmt.Printf("File was converted successfully! Output - %s\n", outputFilename)
}

func exampleConvertOggToMP4() {
	// we don't know anything about the external library
	filename := "some-file.ogg"
	outputFilename := "some-file.mp4"
	wantFormat := "mp4"

	result := facade.NewVideoConverter().Convert(filename, outputFilename, wantFormat)
	if !result {
		panic("file could not be saved")
	}

	fmt.Printf("File was converted successfully! Output - %s\n", outputFilename)
}
