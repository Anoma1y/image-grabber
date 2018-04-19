package main

import (
    "net/http"
    "time"
    "log"
    "github.com/PuerkitoBio/goquery"
    "fmt"
    "strings"
    "os"
    "io"
)

func main() {
    var tag string = "sousouman"
    go grabber(tag)
    grabber(tag)
}

func grabber(tag string) {

    var checkIsEmpty bool = true
    page := 1

    for checkIsEmpty {

        var url string = fmt.Sprintf("https://yande.re/post?page=%d&tags=%s", page, tag)

        tr := &http.Transport{
            MaxIdleConns:       10,
            IdleConnTimeout:    30 * time.Second,
            DisableCompression: true,
        }
        client := &http.Client{Transport: tr}
        res, err := client.Get(url)

        if err != nil {
            log.Fatal(err)
        }

        defer res.Body.Close()
        if res.StatusCode != 200 {
            log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
        }

        doc, err := goquery.NewDocumentFromReader(res.Body)
        if err != nil {
            log.Fatal(err)
        }

        checkIsEmpty = doc.Find("#post-list-posts").Is("ul")
        fmt.Println(checkIsEmpty)

        if checkIsEmpty {
            doc.Find("#post-list-posts li").Each(func(i int, s *goquery.Selection) {
                if val, ok := s.Find("a.directlink").Attr("href"); !ok {
                    fmt.Printf("Nope\n")
                } else {

                   inputImgName := strings.Split(val, "/")
                   imgURL := inputImgName[len(inputImgName) - 1]
                   sourceURL := fmt.Sprintf("D:/test/%s", imgURL)

                   response, e := http.Get(val)

                   if e != nil {
                      log.Fatal(e)
                   }

                   defer response.Body.Close()

                   file, err := os.Create(sourceURL)

                   if err != nil {
                      log.Fatal(err)
                   }

                   _, err = io.Copy(file, response.Body)
                   if err != nil {
                      log.Fatal(err)
                   }

                   file.Close()
                }
            })
        }

        page++
    }

}


func ExampleScrape() {



    //doc.Find("#post-list-posts li").Each(func(i int, s *goquery.Selection) {
    // if val, ok := s.Find("a.directlink").Attr("href"); !ok {
    //   fmt.Printf("Nope\n")
    // } else {
    //     fmt.Printf("%d - %s\n", i, val)
	//
    //    response, e := http.Get(val)
	//
    //    if e != nil {
    //       log.Fatal(e)
    //    }
	//
    //    defer response.Body.Close()
	//
    //    file, err := os.Create(url)
	//
    //    if err != nil {
    //       log.Fatal(err)
    //    }
	//
    //    _, err = io.Copy(file, response.Body)
    //    if err != nil {
    //       log.Fatal(err)
    //    }
    //    fmt.Println("Success")
    //    file.Close()
    //  }
    //})
  }
  