package main

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var file = "clientchain.pem"

func camelCasifinatifier(s string) (r string) {
	for _, i := range strings.Split(s, " ") {
		r += strings.Title(i)
	}
	return
}

func main() {

	// Read the file into "bytes".
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// In the following loop we need access to "bytes", so we define
	// "block" in the outer scope as well.
	var block *pem.Block

	for {
		// Check if there are more bytes left that we need to process.
		if len(bytes) == 0 {
			log.Println("Done...")
			break
		}

		// Decode pem block, and return bytes left to work trough.
		block, bytes = pem.Decode(bytes)
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			log.Fatal(err)
		}

		outputFilename := camelCasifinatifier(cert.Subject.CommonName) + ".pem"

		f, err := os.Create("output/" + outputFilename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		log.Printf("Writing: %v", f.Name())
		pem.Encode(f, block)
		if err != nil {
			log.Fatal(err)
		}
	}
}
