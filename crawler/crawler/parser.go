package crawler

import (
    "io"
    "code.google.com/p/go.net/html"
)


type SimpleParser struct {}


func (parser SimpleParser) ExtractUrls(body io.Reader) (urls []string, err error) {
    urls = make([]string, 10)

    doc, err := html.Parse(body)
    if err != nil {
        return nil, err
    }

    var f func (*html.Node)
    f = func (n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, attr := range n.Attr {
                if attr.Key == "href" {
                    urls = append(urls, attr.Val)
                    break
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }

    f(doc)

    return urls, nil }
