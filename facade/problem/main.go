package main

import (
	"fmt"

	"github.com/arthurshafikov/patterns/facade/problem/lib"
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
	filename := "some-file.mp4"
	outputFilename := "some-file.ogg"
	wantFormat := "ogg"

	file := lib.NewVideoFile(filename)
	sourceCodec := lib.NewCodecFactory().Extract(file)

	var destinationCodec any
	switch wantFormat {
	case "ogg":
		destinationCodec = lib.NewOggCompressionCodec()
	case "mp4":
		destinationCodec = lib.NewMPEG4CompressionCodec()
	}

	bitrateReaderWriter := lib.NewBitrateReaderWriter()

	buffer := bitrateReaderWriter.Read(filename, sourceCodec)
	outputFile := bitrateReaderWriter.Convert(buffer, destinationCodec, outputFilename)

	audioMixer := lib.NewAudioMixer()
	outputFile = audioMixer.FixAudioInVideoFile(outputFile)

	result := bitrateReaderWriter.Save(outputFile)
	if !result {
		panic("file could not be saved")
	}

	fmt.Printf("File was converted successfully! Output - %s\n", outputFile.Filename)
}

func exampleConvertOggToMP4() {
	filename := "some-file.ogg"
	outputFilename := "some-file.mp4"
	wantFormat := "mp4"

	file := lib.NewVideoFile(filename)
	sourceCodec := lib.NewCodecFactory().Extract(file)

	var destinationCodec any
	switch wantFormat {
	case "ogg":
		destinationCodec = lib.NewOggCompressionCodec()
	case "mp4":
		destinationCodec = lib.NewMPEG4CompressionCodec()
	}

	bitrateReaderWriter := lib.NewBitrateReaderWriter()

	buffer := bitrateReaderWriter.Read(filename, sourceCodec)
	outputFile := bitrateReaderWriter.Convert(buffer, destinationCodec, outputFilename)

	audioMixer := lib.NewAudioMixer()
	outputFile = audioMixer.FixAudioInVideoFile(outputFile)

	result := bitrateReaderWriter.Save(outputFile)
	if !result {
		panic("file could not be saved")
	}

	fmt.Printf("File was converted successfully! Output - %s\n", outputFile.Filename)
}
