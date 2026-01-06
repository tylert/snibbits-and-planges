/*usr/bin/env go run "$0" "$@"; exit;*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	FetchLatestGo()
}

func FetchLatestGo() {
	link, checksum := FindLatestGo("linux", "amd64")
	fmt.Println(link)
	fmt.Println(checksum)

	FetchFile(link)
	// XXX FIXME TODO  Check the downloaded file against the expected checksum
	// XXX FIXME TODO  Unzip the archive and provide a way to set the location
}

func FindLatestGo(os string, arch string) (string, string) {
	// Fetch the webby stuff
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get("https://go.dev/dl")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK { // 200
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Stop after finding the first match (should be the newest release)
	link := ""
	checksum := ""
	doc.Find("tr").EachWithBreak(func(i int, s *goquery.Selection) bool {
		a := s.Find("td.filename a")
		tt := s.Find("td tt").Text()
		href, ok := a.Attr("href")
		if ok {
			if strings.Contains(href, fmt.Sprintf("%s-%s.tar.gz", os, arch)) {
				link = fmt.Sprintf("https://go.dev%s", href)
				checksum = tt
				return false
			}
		}
		return true
	})

	return link, checksum
}

func FetchFile(url string) {
	// Connect to the http(s) endpoint
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK { // 200
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	// Create the file locally
	filepath := path.Base(res.Request.URL.String())
	out, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Write to the buffer in 32 kB chunks and count downloaded bytes
	buffer := make([]byte, 32*1024)
	total := res.ContentLength
	var downloaded int64 = 0
	for {
		n, err := res.Body.Read(buffer)
		if n > 0 {
			_, writeErr := out.Write(buffer[:n])
			if writeErr != nil {
				log.Fatal(err)
			}
			downloaded += int64(n)
			fmt.Printf("\rFetching... %d%% complete", 100*downloaded/total)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("\n")
}
