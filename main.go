package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	for i := 0; i < 2000; i++ {
		file := fmt.Sprintf("0x02E%s.png", pad(i))

		log.Print("Downloading: ", file)
		if err := DownloadFile(fmt.Sprintf("https://d1u1mce87gyfbn.cloudfront.net/game/heroes/small/%s", file), path.Join("./icons", file)); err != nil {
			log.Print("Error: ", err)
		}
	}
}

func pad(num int) string {
	return strings.ToUpper(fmt.Sprintf("%013x", num))
}

func DownloadFile(url string, filepath string) error {
	// Create the file

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return errors.New("not found")
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}