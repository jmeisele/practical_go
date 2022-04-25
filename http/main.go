package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello from main.go in http module")
	resp, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		// Fatal will print to stdout and exit the program like sys.exist() in Python
		// main should be the only function that forcefully exits the program
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("bad status: %s", resp.Status)
	}

	fmt.Println("content type:" , resp.Header.Get("Content-Type"))
	fmt.Println("body", resp.Body)
	// io.Copy(os.Stdout, resp.Body)
	var ir IPResponse
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&ir); err != nil {
		log.Fatal(err)
	}
	fmt.Println(ir.Origin)

}

type IPResponse struct {
	Origin string  // Must be exported to work with encoding/json
}