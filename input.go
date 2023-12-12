package aoc

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func GetInput(year, day int) error {
	inputPath := filepath.Join(YearsDir, fmt.Sprintf("%d/%02d.txt", year, day))
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		fHandle, err := os.Create(inputPath)
		if err != nil {
			return fmt.Errorf("%w: failed to create file: %s", err, inputPath)
		}
		defer fHandle.Close()

		url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		env := os.Getenv("AOC_SESSION")
		if env == "" {
			return errors.New("no session token for AOC found")
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: env})

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		readBytes, err := io.ReadAll(resp.Body)
		written, err := fHandle.Write(readBytes)
		if err != nil {
			return err
		}
		fmt.Printf("bytes written: %d, content length: %d", written, resp.ContentLength)
		return nil
	}
	return nil
}

func LoadInput(path string) ([]byte, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return fileBytes, nil
}
