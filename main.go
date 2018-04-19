package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "log"
    "net/http"
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
    res, err := http.Get("http://metalsucks.net")
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
  
    // Find the review items
    doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
      // For each item found, get the band and title
      band := s.Find("a").Text()
      title := s.Find("i").Text()
      fmt.Printf("Review %d: %s - %s\n", i, band, title)
    })
  }
  