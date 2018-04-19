package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "log"
    "net/http"
    "time"
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
    res, err := client.Get("https://yande.re/post?page=1&tags=sousouman")
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

    //Check list is empty
    fuck := doc.Find("#post-list-posts").Is("ul")
    fmt.Println(fuck)


    //doc.Find("#post-list-posts li").Each(func(i int, s *goquery.Selection) {
    //  // For each item found, get the band and title
    //  if val, ok := s.Find("a.directlink").Attr("href"); !ok {
    //    fmt.Printf("Nope\n")
    //  } else {
    //      fmt.Printf("%d - %s\n", i, val)




        //response, e := http.Get(val)
		//
        //if e != nil {
        //    log.Fatal(e)
        //}
		//
        //defer response.Body.Close()
		//
        //img := strings.Split(val, ".")
        //imgType := img[len(img) - 1]
        //imgName := strings.Split(img[2], "/")
        //url := fmt.Sprintf("D:/test/%s.%s", imgName[len(imgName) - 1], imgType)
		//
        //file, err := os.Create(url)
		//
        //if err != nil {
        //    log.Fatal(err)
        //}
		//
        //_, err = io.Copy(file, response.Body)
        //if err != nil {
        //    log.Fatal(err)
        //}
        //fmt.Println("Success")
        //file.Close()
      //}
    //})
  }
  