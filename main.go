package main

import (
    "net/http"
    "log"
    "github.com/PuerkitoBio/goquery"
    "fmt"
    "strings"
    "os"
    "io"
)

func main() {
    var tag = []string{"witch", "dress", "touhou"}
    grabber(tag)
}

func grabber(tag []string) {

    var (
        checkIsEmpty bool = true
        page int = 1
    )

    var tags = getTagsList(tag)

    for checkIsEmpty {

       var url string = fmt.Sprintf("https://yande.re/post?page=%d&tags=%s", page, tags)

       res, err := http.Get(url)

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

       if checkIsEmpty {
          doc.Find("#post-list-posts li").Each(func(i int, s *goquery.Selection) {
              if val, ok := s.Find("a.directlink").Attr("href"); !ok {
                  fmt.Printf("Nope\n")
              } else {
                  download(val)
              }
          })
       }

       page++
    }
}

func getImageName(value string) string {
    inputImgName := strings.Split(value, "/")
    imgURL := inputImgName[len(inputImgName) - 1]
    return fmt.Sprintf("D:/test/%s", imgURL)
}

func getTagsList(tag []string) string {

    var tags string = ""

    if len(tag) > 1 {
        for i := 0; i < len(tag); i++  {
            if i == 0 {
                tags += tag[i]
                continue
            }
            tags += "+" + tag[i]
        }
    } else {
        tags += tag[0]
    }
    return tags
}

func download(value string) {

    var sourceURL string = getImageName(value)

    response, e := http.Get(value)

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