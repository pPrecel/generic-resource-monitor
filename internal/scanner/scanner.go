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

func ExtractVal(filename string, line, elem int) (float64, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return 0.0, err
	}
	defer reader.Close()
	scanner := bufio.NewScanner(reader)

	for lineNumber := 0; scanner.Scan(); lineNumber++ {
		if lineNumber != line {
			continue
		}

		lineText := strings.Fields(scanner.Text())
		if len(lineText) < elem {
			return 0.0, fmt.Errorf(errorFormat, line, elem)
		}

		float, err := strconv.ParseFloat(lineText[elem], 64)
		if err != nil {
			return 0.0, fmt.Errorf(parseErrorFormat, lineText[elem])
		}

		return float, nil
	}

	return 0.0, fmt.Errorf(errorFormat, line, elem)
}
