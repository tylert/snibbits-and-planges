/*usr/bin/env go run "$0" "$@"; exit;*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

func main() {
	// Fetch a remote file
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get("https://somethingsomethingsomething")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK { // 200
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	// Create the file locally
	filepath := path.Base(res.Request.URL.String())
	out, err := os.Create(filepath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer out.Close()

	// Write to the buffer in 32 kB chunks
	buffer := make([]byte, 32*1024)
	totalBytes := res.ContentLength
	var downloadedBytes int64 = 0
	for {
		n, err := res.Body.Read(buffer)
		if n > 0 {
			_, writeErr := out.Write(buffer[:n])
			if writeErr != nil {
				log.Fatalf("Error: %v", err)
			}
			downloadedBytes += int64(n)
			fmt.Printf("\rFetching... %d%% complete", 100*downloadedBytes/totalBytes)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
	fmt.Printf("\n")
}
