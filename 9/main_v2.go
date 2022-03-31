package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	bufSize = 1024 * 8
)

func Wget(fileName string, urlPath string) error {
	tr := new(http.Transport)
	client := &http.Client{Transport: tr}
	resp, err := client.Get(urlPath)

	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bufferedWriter := bufio.NewWriterSize(file, bufSize)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(bufferedWriter, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatal("Please, enter url and fileName:\n ")
	}

	fileName := args[0]
	urlPath := args[1]
	if err := Wget(fileName, urlPath); err != nil {
		log.Printf("Error: %s", err)
	}

}

