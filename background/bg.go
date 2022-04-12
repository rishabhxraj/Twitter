package background

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func GetBackground() error {
	url := "https://source.unsplash.com/random/1600x1600/?art,nature"
	response, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "failed to get image")
	}
	defer response.Body.Close()
	file, err := os.Create(filepath.Join("static", "img", "bg.png"))
	if err != nil {
		return errors.Wrap(err, "failed to create image")
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return errors.Wrap(err, "failed to copy image")
	}
	fmt.Println("Successfully fetched background image")
	return nil
}
