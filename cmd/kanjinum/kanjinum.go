package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"

	"github.com/yakamon/shconf/lib/kanjinum"
)

func main() {
	nums, err := ParseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, n := range nums {
		s := kanjinum.NumToKanji(n)
		fmt.Println(s)
	}
}

// ParseArgs parses args from command line
func ParseArgs() (nums []*big.Int, err error) {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return nil, fmt.Errorf("Failed to read input: %v", err)
	}
	if fileInfo.Mode()&os.ModeNamedPipe != 0 {
		defer os.Stdin.Close()

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if i, ok := new(big.Int).SetString(scanner.Text(), 10); ok {
				nums = append(nums, i)
			} else {
				return nil, fmt.Errorf("Faild to read input: %v", i)
			}
		}
		return nums, nil
	}

	cmdName := os.Args[0]
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <number>\n", cmdName)
	}

	if i, ok := new(big.Int).SetString(os.Args[1], 10); ok {
		nums = append(nums, i)
	} else {
		return nil, fmt.Errorf("Faild to read input: %v", i)
	}
	return nums, nil
}
