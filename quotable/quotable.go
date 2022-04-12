package quotable

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type Quote struct {
	ID           string   `json:"_id"`
	Tags         []string `json:"tags"`
	Content      string   `json:"content"`
	Author       string   `json:"author"`
	AuthorSlug   string   `json:"authorSlug"`
	Length       int      `json:"length"`
	DateAdded    string   `json:"dateAdded"`
	DateModified string   `json:"dateModified"`
}

// GetQuotes fetches quotes from the API and return Quote struct
func GetQuotes() (*Quote, error) {
	quotableURL := "https://api.quotable.io/random?maxLength=150"
	response, err := http.Get(quotableURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to do GET request")
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(err, "got invalid status code: %d", response.StatusCode)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read body")
	}
	var result Quote
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		return nil, errors.Wrap(err, "failed to unmarshal JSON")
	}
	fmt.Println("Sucessfully feched quote")
	return &result, nil
}
