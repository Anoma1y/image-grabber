package main

import (
	"net/http"
	"log"
	"fmt"
	"strings"
	"os"
	"io"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

const (
	PATH string = "D:/test/"
)

var (
	URL string = "https://yande.re/post?page="
	TAG = []string{"nijisanji"}
	DOWNLOAD_LINK []string
)

func main() {
	var tags = getTagsList(TAG)
	var countPage = getCountPage(tags)
	for i := 1; i <= countPage; i++ {
		grab(i, tags)
	}

}

func grab(page int, tags string) {

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

	doc.Find("#post-list-posts li").Each(func(i int, s *goquery.Selection) {
		if val, ok := s.Find("a.directlink").Attr("href"); !ok {
			fmt.Printf("Nope\n")
		} else {
			DOWNLOAD_LINK = append(DOWNLOAD_LINK, val)
		}
	})
}

func getImageName(value string) string {
	var sep string = "/"

	inputImgName := strings.Split(value, "/")
	imgURL := inputImgName[len(inputImgName)-1]

	if string(PATH[len(PATH)-1]) == "/" {
		sep = ""
	}

	return fmt.Sprintf("%s%s%s", PATH, sep, imgURL)
}

func getTagsList(tag []string) string {

	var tags string = ""

	if len(tag) == 1 {
		return tag[0]
	}

	for i := 0; i < len(tag); i++ {
		if i == 0 {
			tags += tag[i]
			continue
		}
		tags += "+" + tag[i]
	}

	return tags
}

func getCountPage(tags string) int {
	var url string = fmt.Sprintf("%s%d&tags=%s", URL, 1, tags)

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

	count, err := strconv.Atoi(doc.Find(".pagination a:nth-last-of-type(2)").Text())
	if err != nil {
		return 0
	}
	return count
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
