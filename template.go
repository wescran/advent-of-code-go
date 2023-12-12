package aoc

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

const YearsDir string = "/Users/wcranston/gitrepos/advent-of-code-go"
const solutionTemplate string = `package year{{.Year}}

import (
	"fmt"
	aoc "github.com/wescran/advent-of-code-go"
)

type Solution{{.Year}}{{printf "%02d" .Day}} struct{}

func (s *Solution{{.Year}}{{printf "%02d" .Day}}) Solve() error{
	data, err := aoc.LoadInput(fmt.Sprintf("%s/%d/%02d.txt", aoc.YearsDir, {{.Year}}, {{.Day}}))
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	return nil
}`

func CreateSolutionFile(year, day int) error {
	temp, err := template.New("solution").Parse(solutionTemplate)
	if err != nil {
		return err
	}
	yearString := fmt.Sprintf("%d", year)
	dayString := fmt.Sprintf("%02d.go", day)

	dirPath := filepath.Join(YearsDir, yearString)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, 0750)
		if err != nil {
			return fmt.Errorf("failed to create dir: %s", dirPath)
		}
	}
	solutionPath := filepath.Join(dirPath, dayString)
	if _, err := os.Stat(solutionPath); os.IsNotExist(err) {
		fHandle, err := os.Create(filepath.Join(dirPath, dayString))
		if err != nil {
			return fmt.Errorf("failed to create solution file: %s", dirPath)
		}
		defer fHandle.Close()

		temp.Execute(fHandle, struct {
			Year int
			Day  int
		}{
			Year: year,
			Day:  day,
		})
		fmt.Printf("created solution file: %s\n", solutionPath)
		return nil
	}
	fmt.Printf("solution file exists, noop: %s\n", solutionPath)
	return nil
}
