package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func DownloadFile(filepath string, url string) (err error) {
	fmt.Println(color.GreenString("-- Acquiring"), filepath)

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
