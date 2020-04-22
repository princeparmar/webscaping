package parser

import (
	"fmt"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/princeparmar/webscaping/outer"
	"golang.org/x/net/html"
)

// Collection finds all the link from the font page.
func Collection(category string, pageno int) ([]Product, error) {
	fmt.Println("downloading collection", category, "pageno", pageno)
	doc, err := outer.Get(fmt.Sprintf("https://www.bridallehengastore.com/collections/%s?page=%d", category, pageno))
	if err != nil {
		return nil, err
	}

	var out []Product
	var wg sync.WaitGroup
	var mx sync.Mutex

	doc.Find("a.grid-link").Each(func(i int, el *goquery.Selection) {
		for _, v := range el.Nodes {
			wg.Add(1)
			func(node *html.Node) {
				defer wg.Done()
				url := FindAttr(*v, "href")
				p, err := FetchProduct(url)
				if err != nil {
					fmt.Println("error in product fetch", err)
					return
				}

				mx.Lock()
				out = append(out, *p)
				mx.Unlock()
			}(v)
		}
	})

	wg.Wait()
	fmt.Println("downloaded collection", category, "pageno", pageno)
	return out, nil

}

func FullCollection(name string) ([]Product, error) {
	var out []Product
	var mx sync.Mutex

	for i := 0; ; i++ {
		tmp, err := Collection(name, i)
		if err != nil {
			return nil, err
		}

		if len(tmp) == 0 {
			return out, nil
		}

		mx.Lock()
		out = append(out, tmp...)
		mx.Unlock()
	}
}
