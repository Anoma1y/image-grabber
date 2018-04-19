package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "log"
    "net/http"
    "time"
    // "io"
)

func main() {
    ExampleScrape()
    // url := "http://i.imgur.com/m1UIjW1.jpg"
    

    // response, e := http.Get(url)

    // if e != nil {
    //     log.Fatal(e)
    // }


    // defer response.Body.Close()

    // file, err := os.Create("/home/developer/Projects/qwerty.jpg")

    // if err != nil {
    //     log.Fatal(err)
    // }

    // _, err = io.Copy(file, response.Body)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // file.Close()
    // fmt.Println("Success!")
}
func ExampleScrape() {
    // Request the HTML page.
    tr := &http.Transport{
      MaxIdleConns:       10,
      IdleConnTimeout:    30 * time.Second,
      DisableCompression: true,
    }
    client := &http.Client{Transport: tr}
    res, err := client.Get("https://lenta.ru/")
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
    doc.Find(".js-main__sidebars .b-yellow-box__wrap .item").Each(func(i int, s *goquery.Selection) {
      // For each item found, get the band and title
      band := s.Find("a").Text()
      fmt.Printf("Review %d: %s\n", i, band)
    })
  }
  