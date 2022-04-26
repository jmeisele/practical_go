package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	sig, err := fileSHA1("http.log.gz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sig)
}

func fileSHA1(fileName string) (string, error) {
	// Read the file
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	
	r, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}
	
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}
	
	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil

}

