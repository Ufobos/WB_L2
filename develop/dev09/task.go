package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

var url, outputFile string

func init() {
	flag.StringVar(&url, "url", "", "URL to download")
	flag.StringVar(&outputFile, "output", "output.html", "File name to save the downloaded content")
}

func main() {
	flag.Parse()
	// Проверяем, что URL действительно передан
	if url == "" {
		log.Fatal("URL is required")
	}

	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching URL %s: %v", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Server return non-200 status: %d %s", response.StatusCode, response.Status)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error writing response to file: %v", err)
	}

	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatalf("Error writing response to file: %v", err)
	}

}
