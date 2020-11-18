package scanner

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	errorFormat      = "can't find value for line: %d, elem: %d"
	parseErrorFormat = "can't parse value: %s as float64"
)

func ExtractFloat(lines []string, line, column int) (float64, error) {
	if len(lines)-1 < line {
		return 0.0, fmt.Errorf(errorFormat, line, column)
	}

	lineText := strings.Fields(lines[line])
	if len(lineText)-1 < column {
		return 0.0, fmt.Errorf(errorFormat, line, column)
	}

	float, err := strconv.ParseFloat(lineText[column], 64)
	if err != nil {
		return 0.0, fmt.Errorf(parseErrorFormat, lineText[column])
	}

	return float, nil
}

func ReadLines(filename string) ([]string, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	scanner := bufio.NewScanner(reader)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
