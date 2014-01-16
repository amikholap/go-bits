package main

import (
    "fmt"
    "log"
    "os"
    "strconv"
    "go-bits/crawler/crawler"
)

func main() {
    start_url := os.Args[1]
    depth, err := strconv.ParseUint(os.Args[2], 10, 64)
    if err != nil {
        log.Fatal(err)
    }

    const dirname = "crawled_documents"
    err = os.Mkdir(dirname, 0755)
    if err != nil && os.IsNotExist(err) {
        log.Fatal(err)
    }

    urls, err := crawler.Crawl(dirname, start_url, depth)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Crawled %d urls\n", len(urls))

    total_len := 0
    for _, v := range urls {
        total_len += len(v)
    }
    
    fmt.Printf("Total size: %dK\n", total_len / 1024)
}
