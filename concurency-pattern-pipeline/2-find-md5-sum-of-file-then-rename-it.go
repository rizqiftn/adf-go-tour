package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.59-pipeline-temp")

func main() {
	log.Println("Start.")
	start := time.Now()

	proceed()

	duration := time.Since(start)
	log.Println("Done in ", duration.Seconds(), "seconds")
}

func proceed() {
	counterTotal := 0
	counterRenamed := 0

	err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fmt.Println("File executed")
		counterTotal++

		// read file
		buf, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// sum int
		sum := fmt.Sprintf("%x", md5.Sum(buf))

		// rename file
		destinationPath := filepath.Join(tempPath, fmt.Sprintf("file-%s-%d.txt", sum, counterRenamed))
		err = os.Rename(path, destinationPath)
		fmt.Println("Path: ", path)
		fmt.Println("Destination Path: ", destinationPath)
		if err != nil {
			return err
		}

		counterRenamed++
		return nil
	})

	if err != nil {
		log.Println("Error: ", err.Error())
	}

	log.Printf("%d/%d files Renamed", counterRenamed, counterTotal)
}
