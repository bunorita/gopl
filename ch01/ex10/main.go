package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		go fetch(url, ch, fmt.Sprintf("out%d.html", i))
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	wg.Wait() // wait for writing file
}

func fetch(url string, ch chan<- string, ofile string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	// write file by another goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		f, err := os.Create(ofile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create file: %s\n", ofile)
			return
		}
		f.Write(b)
		f.Close()
	}()

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, len(b), url)
}
