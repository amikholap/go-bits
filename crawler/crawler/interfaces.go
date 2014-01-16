package crawler


type Fetcher interface {
    // Fetch returns the body of URL and a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}


type Parser interface {
    // ExtractUrls returns all links found in the document's body
    ExtractUrls(body string) (urls []string, err error)
}
