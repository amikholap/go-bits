package crawler

import (
    "bytes"
    "io/ioutil"
    "net/http"
)


type SimpleFetcher struct {}


func (fetcher SimpleFetcher) Fetch(url string) (string, []string, error) {
    var parser SimpleParser

    resp, err := http.Get(url)
    if err != nil {
        return "", nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", nil, err
    }

    urls, _ := parser.ExtractUrls(bytes.NewReader(body))

    return string(body), urls, nil
}
