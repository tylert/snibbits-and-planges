/*usr/bin/env go run "$0" "$@"; exit;*/

/* Download the latest version of the Go compiler via HTTP, verify the
   checksum and extract it so we can start using it. */

/* XXX FIXME TODO
- Fix executable bits getting squashed when extracting the tarball
- Allow fetching other release versions
- Waaaaay better handling of desired locations and filenames
- Check if we already have the desired version of the compiler installed
- Better handling when the tarball has already been downloaded
- Find a better way than printf to report download progress
- Throw errors from functions and handle fatal calls from main
*/

package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	//"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	FetchLatestGo()
}

func FetchLatestGo() {
	// Figure out the link to the latest version (and the expected checksum)
	url, checksum := FindGoReleaseTarballLink(runtime.GOOS, runtime.GOARCH)
	fmt.Println(url)
	fmt.Println(checksum)

	// Download the latest version from the provided link
	dest := HTTPGetFile(url, "")

	// Verify the checksum against the one we are expecting
	hash := SHA256File(dest)
	fmt.Println(hash)

	ExtractTarGzFileToDisk(dest)
}

func FindGoReleaseTarballLink(os string, arch string) (string, string) {
	// Do a bit of web-scraping to get a link to the latest version tarball and its checksum
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
		log.Fatal(fmt.Sprintf("Can't find a tarball link for os=%s, arch=%s.", os, arch))
	}

	return link, checksum
}

func HTTPGetFile(url string, dest string) string {
	// Connect to the desired endpoint
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK { // 200
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	// Figure out what to name the file and where to put it
	fullpath := dest
	if dest == "" {
		fullpath = path.Base(res.Request.URL.String())
	}

	/*
		if _, err := os.Stat(dest); errors.Is(err, os.ErrNoExist) {
			// it does not exist
		}
		if info, err := os.Stat(dest); err != nil {
			// it does exist but now we can figure out if it's a dir or a file
		}
		if info.IsDir() {
		}

		dir := path.Dir(dest)
		file := path.Base(dest)
		fullpath := filepath.Join(dir, file)
	*/

	// Create the local file (fail if the file already exists)
	fd, err := os.OpenFile(fullpath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	// Write to the buffer in 32 kB chunks and count downloaded bytes
	buffer := make([]byte, 32*1024)
	total := res.ContentLength
	var downloaded int64 = 0
	for {
		n, err := res.Body.Read(buffer)
		if n > 0 {
			_, writeErr := fd.Write(buffer[:n])
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

func SHA256File(file string) string {
	fd, err := os.OpenFile(file, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, fd); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func ExtractTarGzFileToDisk(file string) {
	fd, err := os.OpenFile(file, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	// https://mustafanafizdurukan.github.io/posts/understanding-io-limitreader

	ungz, err := gzip.NewReader(fd)
	// foo := io.LimitReader(ungz, n)
	// bar := bufio.NewReader(foo)
	if err != nil {
		log.Fatal(err)
	}
	defer ungz.Close()

	unt := tar.NewReader(ungz)
	// baz := io.LimitReader(unt, n)
	// quux := bufio.NewReader(baz)
	var header *tar.Header
	for header, err = unt.Next(); err == nil; header, err = unt.Next() {
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(header.Name, 0755); err != nil {
				log.Fatal(err)
			}
		case tar.TypeReg:
			// Skip any weird muckOS resource files
			if !strings.HasPrefix(header.Name, "._") {
				out, err := os.OpenFile(header.Name, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
				if err != nil {
					log.Fatal(err)
				}
				if _, err := io.Copy(out, unt); err != nil {
					out.Close()
					log.Fatal(err)
				}
				if err := out.Close(); err != nil {
					log.Fatal(err)
				}
			}
		default: // hope that nobody is expecting symlinks or other non-regular file types to survive
			log.Fatal("You should never see this error")
		}
	}
	if err != io.EOF {
		log.Fatal("There was some kind of unexplained error")
	}
}
