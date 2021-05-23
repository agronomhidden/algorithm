package tester

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"time"
)

type Interface interface {
	Run(input []string) string
}

type Option struct {
	MaxIteration  int
	PreventOutput bool
	Debug bool
}

func Run(testDir string, task Interface, opt Option)  {
	i := 0
	for {
		inFile  := path.Join(testDir, fmt.Sprintf("test.%d.in", i))
		outFile := path.Join(testDir, fmt.Sprintf("test.%d.out", i))
		input, errIn := readFileLineByLine(inFile)
		output, errOut := readFileLineByLine(outFile)
		if errIn != nil || errOut != nil {
			if opt.Debug {
				fmt.Println(errIn)
				fmt.Println(errOut)
			}
			break
		}
		intime := time.Now()
		result := task.Run(input)
		runtime := time.Now().Sub(intime).Microseconds()
		fmt.Printf("Test %d: %t [%vμs]\n", i, result == output[0], runtime)
		if result != output[0] {
			if opt.PreventOutput == false {
				fmt.Printf("\texp:\t%s\n", output[0])
				fmt.Printf("\tact:\t%s\n", result)
			}
		}
		if opt.MaxIteration > 0 && opt.MaxIteration == i {
			break
		}
		i++
	}
}

func readFileLineByLine(filePath string) (output []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	err = scanner.Err()
	return
}




