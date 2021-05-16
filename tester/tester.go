package tester

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"sync"
	"time"
)

type Interface interface {
	Run(input []string) string
}

func Run(testDir string, task Interface)  {
	i := 0
	var wg sync.WaitGroup
	for {
		inFile  := path.Join(testDir, fmt.Sprintf("test.%d.in", i))
		outFile := path.Join(testDir, fmt.Sprintf("test.%d.out", i))
		input, errIn := readFileLineByLine(inFile)
		output, errOut := readFileLineByLine(outFile)
		if errIn != nil || errOut != nil {
			break
		}

		wg.Add(1)
		go func(input,  output []string, i int) {
			intime := time.Now()
			result := task.Run(input)
			runtime := time.Now().Sub(intime).Microseconds()
			fmt.Printf("Test %d: %t [%vμs]\n", i, result == output[0], runtime)
			wg.Done()
		}(input, output, i)
		i++
	}
	wg.Wait()
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




