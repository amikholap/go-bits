package crawler

import (
    "net/url"
    "os"
    "path/filepath"
    "strings"
    "sync"
)


func saveDocument(output_dir, rawurl string, content []byte) (err error) {
    u, err := url.Parse(rawurl)

    host_dir := filepath.Join(output_dir, u.Host)
    err = os.Mkdir(host_dir, 0755)
    if err != nil && os.IsNotExist(err) {
        return
    }

    filename := u.Path
    if len(filename) == 0 {
        filename = "/"
    }
    filename = strings.Replace(filename, "/", ":", -1)
    path := filepath.Join(host_dir, filename)
    file, err := os.Create(path)
    if err != nil {
        return
    }
    _, err = file.Write(content)
    if err != nil {
        return
    }
}


func Crawl(output_dir string, initial_url string, depth uint64) map[string]string {
    const nGos = 100
    var fetcher SimpleFetcher
    limiter := make(chan int, nGos)
    var wg sync.WaitGroup
    visited_urls := make(map[string]string)

    var DoCrawl func (string, uint64)
    DoCrawl = func (url string, depth uint64) {
        defer wg.Done()
        defer func() { <- limiter }()

        limiter <- 0

        if depth <= 0 {
            return
        }

        body, urls, err := fetcher.Fetch(url)

        if err != nil {
            return
        } else {
            saveDocument(output_dir, url, []byte(body))
            visited_urls[url] = body
        }

        for _, url := range urls {
            if _, present := visited_urls[url]; !present {
                wg.Add(1)
                go DoCrawl(url, depth - 1)
            }
        }
    };

    wg.Add(1)
    go DoCrawl(initial_url, depth)
    wg.Wait()

    return visited_urls
}
