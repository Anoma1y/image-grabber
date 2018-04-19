package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "log"
    "net/http"
    "time"
    "io"
    "os"
)

func main() {
    ExampleScrape()
}
func ExampleScrape() {
    // Request the HTML page.
    tr := &http.Transport{
      MaxIdleConns:       10,
      IdleConnTimeout:    30 * time.Second,
      DisableCompression: true,
    }
    client := &http.Client{Transport: tr}
    res, err := client.Get("https://lenta.ru/rubrics/world/")
    if err != nil {
      log.Fatal(err)
    }
    defer res.Body.Close()
    if res.StatusCode != 200 {
      log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
    }
  
    // Load the HTML document
    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Printf("Start\n")
    // Find the review items
    doc.Find(".article a.picture").Each(func(i int, s *goquery.Selection) {
      // For each item found, get the band and title
      if val, ok := s.Find("img").Attr("src"); !ok {
        fmt.Printf("Nope\n")
      } else {

        image := fmt.Sprintf("/home/developer/Projects/image%d.jpg", i)
        response, e := http.Get(val)
    
        if e != nil {
            log.Fatal(e)
        }
    
    
        defer response.Body.Close()

        file, err := os.Create(image)

        if err != nil {
            log.Fatal(err)
        }

        _, err = io.Copy(file, response.Body)
        if err != nil {
            log.Fatal(err)
        }
        file.Close()
        fmt.Println("Success!")
      }
    })
  }
  