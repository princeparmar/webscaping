package parser

import (
	"fmt"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/princeparmar/webscaping/outer"
	"golang.org/x/net/html"
)

// Collection finds all the link from the font page.
func Collection(category string, pageno int) (bool, error) {
	fmt.Println("downloading collection", category, "pageno", pageno)
	doc, err := outer.Get(fmt.Sprintf("https://www.bridallehengastore.com/collections/%s?page=%d", category, pageno))
	if err != nil {
		return false, err
	}

	var wg sync.WaitGroup
	var shouldContinue bool
	doc.Find("a.grid-link").Each(func(i int, el *goquery.Selection) {
		shouldContinue = true
		for _, v := range el.Nodes {
			wg.Add(1)
			go func(node *html.Node) {
				defer wg.Done()
				url := FindAttr(*node, "href")
				p, err := FetchProduct(url)
				if err != nil {
					fmt.Println("error in product fetch", err)
					return
				}

				err = p.Save()
				if err != nil {
					fmt.Println("error in product save", err)
				}
			}(v)
		}
	})

	wg.Wait()
	fmt.Println("downloaded collection", category, "pageno", pageno)
	return shouldContinue, nil

}

func FullCollection(name string) error {

	for i := 1; ; i++ {
		shouldContinue, err := Collection(name, i)
		if err != nil {
			return err
		}

		if !shouldContinue {
			return nil
		}

	}
}
