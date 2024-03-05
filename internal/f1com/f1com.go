package f1com

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type F1com struct {
}

func openUrl(u string) (doc *goquery.Document, err error) {
	res, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err = goquery.NewDocumentFromReader(res.Body)

	return
}
