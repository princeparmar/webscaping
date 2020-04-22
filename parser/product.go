package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/princeparmar/webscaping/outer"
)

type Product struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Vendor      string     `json:"vendor"`
	Handle      string     `json:"handle"`
	Type        string     `json:"type"`
	Price       int        `json:"price"`
	PriceMax    int        `json:"price_max"`
	PriceMin    int        `json:"price_min"`
	Description string     `json:"description"`
	Media       []Media    `json:"media"`
	Variants    []Variants `json:"variants"`
}

type Variants struct {
	SKU    string `json:"sku"`
	Weight int    `json:"weight"`
}

type Media struct {
	Alt       string `json:"alt"`
	Src       string `json:"src"`
	LocalPath string `json:"local_path"`
}

func FetchProduct(url string) (*Product, error) {
	fmt.Println("downloading", url)
	doc, err := outer.Get("https://www.bridallehengastore.com" + url)
	if err != nil {
		return nil, err
	}

	product := new(Product)

	//grid-link__title
	rawProduct := doc.Find(`#ProductJson-product-template`).Text()
	err = json.Unmarshal([]byte(rawProduct), product)
	if err != nil {
		return nil, err
	}

	err = product.LoadMedia()
	fmt.Println("downloaded", url)

	return product, err
}

func (p *Product) LoadMedia() error {

	basePath := strings.ReplaceAll(p.Type, " ", "")
	for i, m := range p.Media {
		var err error
		m.LocalPath, err = outer.Download(m.Src, path.Join("output", "images", basePath))
		if err != nil {
			return err
		}
		p.Media[i] = m
	}

	return nil
}

func (p Product) Save() error {
	basePath := strings.ReplaceAll(p.Type, " ", "")
	basePath = path.Join("output", "data", basePath)
	err := os.MkdirAll(basePath, 0700)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return outer.SaveFile(path.Join(basePath, fmt.Sprintf("%d", p.ID)), p)
}
