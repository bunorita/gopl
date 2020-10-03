package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Use io.Copy instead of ioutil.ReadAll
// ref: https://haisum.github.io/2017/09/11/golang-ioutil-readall/
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying %s %v\n", url, err)
			os.Exit(1)
		}
	}
}
