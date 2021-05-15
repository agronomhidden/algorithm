package tester

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

type Interface interface {
	Run(input []string) string
}

func Run(testDir string, task Interface)  {
	i := 0
	for {
		inFile  := path.Join(testDir, fmt.Sprintf("test.%d.in", i))
		outFile := path.Join(testDir, fmt.Sprintf("test.%d.out", i))

		input, errIn := readFileLineByLine(inFile)
		output, errOut := readFileLineByLine(outFile)
		if errIn != nil || errOut != nil {
			return
		}
		fmt.Println(input, output)
	}
}

func readFileLineByLine(filePath string) (output []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	err = scanner.Err()
	return
}




