/*usr/bin/env go run "$0" "$@"; exit;*/

/* Download the latest version of the Go compiler suite via HTTP, verify the
   checksum and extract it so we can start using it. */

/* XXX FIXME TODO
- Waaaaay better checking of the desired dest location and filename
- Allow picking os and arch to fetch
- Check if the latest version is already installed so we can skip unnecessary downloading
- Unzip the archive somewhere in the $PATH like "/usr/local/"
- Find a better way than printf to report download progress
- Throw errors from functions and handle fatal calls from main
*/

package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	//"path"
	//"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	FetchLatestGo()
}

func FetchLatestGo() {
	// Figure out the link to the latest version (and the expected checksum)
	url, checksum := DetermineLatestGo("linux", "amd64")
	fmt.Println(url)
	fmt.Println(checksum)

	// Download the latest version from the provided link
	dest := FetchFile(url, "/tmp/foo.tar.gz")

	// Verify the checksum against the one we are expecting
	hash := HashFileSHA256(dest)
	fmt.Println(hash)
}

func DetermineLatestGo(os string, arch string) (string, string) {
	// Do a bit of web-scraping to determine the latest version checksum
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

	// We don't know how to download nothing
	if link == "" {
		log.Fatal(fmt.Sprintf("Can't find an archive link for os=%s, arch=%s.", os, arch))
	}

	return link, checksum
}

func FetchFile(url string, dest string) string {
	// Connect to the desired endpoint
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK { // 200
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	/*
	     dir := path.Dir(dest)
	     file := path.Base(dest)
	     if strings.HasSuffix(dest, "/") {
	       dir = dest
	   		file = path.Base(res.Request.URL.String())
	     }
	     fmt.Println(dir)
	     fmt.Println(file)
	   	fullpath := filepath.Join(dir, file)
	*/
	fullpath := dest

	// Create the local file
	fh, err := os.OpenFile(fullpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	// Write to the buffer in 32 kB chunks and count downloaded bytes
	buffer := make([]byte, 32*1024)
	total := res.ContentLength
	var downloaded int64 = 0
	for {
		n, err := res.Body.Read(buffer)
		if n > 0 {
			_, writeErr := fh.Write(buffer[:n])
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

	return fullpath
}

func HashFileSHA256(file string) string {
	fh, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, fh); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}
