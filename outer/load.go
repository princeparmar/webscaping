package outer

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

// Get loads html pages.
func Get(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "Get error from source")
	}
	defer resp.Body.Close()

	return goquery.NewDocumentFromReader(resp.Body)
}
