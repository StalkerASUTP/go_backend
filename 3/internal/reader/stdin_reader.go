package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StdinReader struct{}

func NewStdinReader() *StdinReader {
	return &StdinReader{}
}

func (r *StdinReader) ReadData() ([]int64, error) {
	const op = "op.StdinReader.ReadData"
	var numbers []int64
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a number (or just Enter to finish): ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			break
		}
		number, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}	
		numbers = append(numbers, number)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return numbers, nil
}

